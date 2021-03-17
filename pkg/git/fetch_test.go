package git

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/BurntSushi/toml"
	gogit "github.com/go-git/go-git/v5"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/sclevine/spec"
	"github.com/stretchr/testify/require"
)

func TestGitCheckout(t *testing.T) {
	spec.Run(t, "Test Describe Image", testGitCheckout)
}

func testGitCheckout(t *testing.T, when spec.G, it spec.S) {
	when("#Fetch", func() {
		outpuBuffer := &bytes.Buffer{}
		fetcher := Fetcher{
			Logger:   log.New(outpuBuffer, "", 0),
			Keychain: fakeGitKeychain{},
		}
		var testDir string
		var metadataDir string

		it.Before(func() {
			var err error
			testDir, err = ioutil.TempDir("", "test-git")
			require.NoError(t, err)

			metadataDir, err = ioutil.TempDir("", "test-git")
			require.NoError(t, err)
		})

		it.After(func() {
			fmt.Println(testDir)
			//require.NoError(t, os.RemoveAll(testDir))
			require.NoError(t, os.RemoveAll(metadataDir))
		})

		testFetch := func(gitUrl, revision string) func() {
			return func() {
				err := fetcher.Fetch(testDir, gitUrl, revision, metadataDir)
				require.NoError(t, err)

				repository, err := gogit.PlainOpenWithOptions(testDir, &gogit.PlainOpenOptions{})
				require.NoError(t, err)

				worktree, err := repository.Worktree()
				require.NoError(t, err)

				status, err := worktree.Status()
				require.NoError(t, err)

				require.True(t, status.IsClean(), "should be clean")

				require.Contains(t, outpuBuffer.String(), fmt.Sprintf("Successfully cloned \"%s\" @ \"%s\"", gitUrl, revision))

				require.FileExists(t, path.Join(metadataDir, "project-metadata.toml"))

				var projectMetadata project
				_, err = toml.DecodeFile(path.Join(metadataDir, "project-metadata.toml"), &projectMetadata)
				require.NoError(t, err)

				require.Equal(t, "git", projectMetadata.Source.Type)
				require.Equal(t, gitUrl, projectMetadata.Source.Metadata.Repository)
				require.Equal(t, revision, projectMetadata.Source.Metadata.Revision)

				h, err := repository.Head()
				require.NoError(t, err)
				require.Equal(t, h.Hash().String(), projectMetadata.Source.Version.Commit)
			}
		}

		it("fetches remote HEAD", testFetch("https://github.com/git-fixtures/basic", "master"))

		it("fetches a branch", testFetch("https://github.com/git-fixtures/basic", "branch"))

		it("fetches a tag", testFetch("https://github.com/git-fixtures/tags", "lightweight-tag"))

		it("fetches a revision", testFetch("https://github.com/git-fixtures/basic", "b029517f6300c2da0f4b651b8642506cd6aaf45d"))

		it("returns error on non-existent ref", func() {
			err := fetcher.Fetch(testDir, "https://github.com/git-fixtures/basic", "doesnotexist", "")
			require.EqualError(t, err, "could not find reference: doesnotexist")
		})

		it("returns error from remote fetch when authentication required", func() {
			err := fetcher.Fetch(testDir, "git@bitbucket.com:org/repo", "main", "")
			require.EqualError(t, err, "error fetching remote: callback returned unsupported credentials type")
		})
	})
}

type fakeGitKeychain struct{}

func (f fakeGitKeychain) Resolve(url string, username_from_url string, allowed_types git2go.CredentialType) (Git2GoCredential, error) {
	return BasicGit2GoAuth{"thisisnotgonnawork", "AtAll"}, nil
}
