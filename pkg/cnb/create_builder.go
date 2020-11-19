package cnb

import (
	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

type RegistryClient interface {
	Fetch(keychain authn.Keychain, repoName string) (v1.Image, string, error)
	Save(keychain authn.Keychain, tag string, image v1.Image) (string, error)
}

type BuildpackRepository interface {
	FindByIdAndVersion(id, version string) (RemoteBuildpackInfo, error)
}

type RemoteBuilderCreator struct {
	RegistryClient RegistryClient
	LifecycleImage string
	KpackVersion   string
}

func (r *RemoteBuilderCreator) CreateBuilder(keychain authn.Keychain, buildpackRepo BuildpackRepository, clusterStack *v1alpha1.ClusterStack, spec v1alpha1.BuilderSpec) (v1alpha1.BuilderRecord, error) {
	buildImage, _, err := r.RegistryClient.Fetch(keychain, clusterStack.Status.BuildImage.LatestImage)
	if err != nil {
		return v1alpha1.BuilderRecord{}, err
	}

	lifecycleImage, _, err := r.RegistryClient.Fetch(keychain, r.LifecycleImage)
	if err != nil {
		return v1alpha1.BuilderRecord{}, err
	}

	builderBldr, err := newBuilderBldr(lifecycleImage, r.KpackVersion)
	if err != nil {
		return v1alpha1.BuilderRecord{}, err
	}

	builderBldr.AddStack(buildImage, clusterStack)

	metadata := make(v1alpha1.MetadataOrder, 0)

	for _, group := range spec.Order {
		buildpacks := make([]RemoteBuildpackRef, 0, len(group.Group))
		groupOrder := make(v1alpha1.BuildpackMetadataList, 0)

		for _, buildpack := range group.Group {
			remoteBuildpack, err := buildpackRepo.FindByIdAndVersion(buildpack.Id, buildpack.Version)
			if err != nil {
				return v1alpha1.BuilderRecord{}, err
			}

			buildpacks = append(buildpacks, remoteBuildpack.Optional(buildpack.Optional))

			groupOrder = append(groupOrder, buildpackMetadata(remoteBuildpack))
		}
		builderBldr.AddGroup(buildpacks...)
		metadata = append(metadata, v1alpha1.MetadataOrderEntry{Group: groupOrder})
	}

	writeableImage, err := builderBldr.WriteableImage()
	if err != nil {
		return v1alpha1.BuilderRecord{}, err
	}

	identifier, err := r.RegistryClient.Save(keychain, spec.Tag, writeableImage)
	if err != nil {
		return v1alpha1.BuilderRecord{}, err
	}

	return v1alpha1.BuilderRecord{
		Image: identifier,
		Stack: v1alpha1.BuildStack{
			RunImage: clusterStack.Status.RunImage.LatestImage,
			ID:       clusterStack.Status.Id,
		},
		Buildpacks: metadata,
	}, nil
}

func buildpackMetadata(buildpack RemoteBuildpackInfo) v1alpha1.BuildpackMetadata {
	groupList := make(v1alpha1.MetadataOrder, 0)
	for _, layer := range buildpack.Layers {
		for _, group := range layer.BuildpackLayerInfo.Order {
			currentGroup := make(v1alpha1.BuildpackMetadataList, 0)
			for _, item := range group.Group {

				currentGroup = append(currentGroup, v1alpha1.BuildpackMetadata{
					Id:       item.Id,
					Version:  item.Version,
				})
			}

			groupList = append(groupList, v1alpha1.MetadataOrderEntry{Group: currentGroup})
		}
	}
	return v1alpha1.BuildpackMetadata{
		Id:       buildpack.BuildpackInfo.Id,
		Version:  buildpack.BuildpackInfo.Version,
		Homepage: buildpack.BuildpackInfo.Homepage,
		Order:    groupList,
	}
}
