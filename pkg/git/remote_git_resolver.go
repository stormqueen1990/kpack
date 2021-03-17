package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	git2go "github.com/libgit2/git2go/v31"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

const defaultRemote = "origin"

type remoteGitResolver struct {
}

func (*remoteGitResolver) Resolve(keychain GitKeychain, sourceConfig v1alpha1.SourceConfig) (v1alpha1.ResolvedSourceConfig, error) {
	dir, err := ioutil.TempDir("", "git-resolve")
	if err != nil {
		return v1alpha1.ResolvedSourceConfig{}, err
	}
	defer os.RemoveAll(dir)

	repository, err := git2go.InitRepository(dir, false)
	if err != nil {
		return v1alpha1.ResolvedSourceConfig{}, err
	}
	defer repository.Free()

	remote, err := repository.Remotes.CreateWithOptions(sourceConfig.Git.URL, &git2go.RemoteCreateOptions{
		Name:  defaultRemote,
		Flags: git2go.RemoteCreateSkipInsteadof,
	})
	if err != nil {
		return v1alpha1.ResolvedSourceConfig{}, err
	}
	defer remote.Free()

	err = remote.ConnectFetch(&git2go.RemoteCallbacks{
		CredentialsCallback: asCredentialCallback(keychain),
		CertificateCheckCallback: func(cert *git2go.Certificate, valid bool, hostname string) git2go.ErrorCode {
			return git2go.ErrorCodeOK

		},
	}, nil, nil)
	if err != nil {
		return v1alpha1.ResolvedSourceConfig{}, err
	}

	references, err := remote.Ls()
	if err != nil {
		return v1alpha1.ResolvedSourceConfig{}, err
	}

	for _, ref := range references {
		for _, format := range refRevParseRules {
			if fmt.Sprintf(format, sourceConfig.Git.Revision) == ref.Name {
				return v1alpha1.ResolvedSourceConfig{
					Git: &v1alpha1.ResolvedGitSource{
						URL:      sourceConfig.Git.URL,
						Revision: ref.Id.String(),
						Type:     sourceType(ref),
						SubPath:  sourceConfig.SubPath,
					},
				}, nil
			}
		}
	}

	return v1alpha1.ResolvedSourceConfig{
		Git: &v1alpha1.ResolvedGitSource{
			URL:      sourceConfig.Git.URL,
			Revision: sourceConfig.Git.Revision,
			Type:     v1alpha1.Commit,
			SubPath:  sourceConfig.SubPath,
		},
	}, nil
}

func sourceType(reference git2go.RemoteHead) v1alpha1.GitSourceKind {
	switch {
	case strings.HasPrefix(reference.Name, "refs/heads"):
		return v1alpha1.Branch
	case strings.HasPrefix(reference.Name, "refs/tags"):
		return v1alpha1.Tag
	default:
		return v1alpha1.Unknown
	}
}

var refRevParseRules = []string{
	"refs/%s",
	"refs/tags/%s",
	"refs/heads/%s",
	"refs/remotes/%s",
	"refs/remotes/%s/HEAD",
}
