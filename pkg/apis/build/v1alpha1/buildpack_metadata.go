package v1alpha1

type BuildpackMetadataList []BuildpackMetadata

type MetadataOrder []MetadataOrderEntry

// +k8s:openapi-gen=true
type MetadataOrderEntry struct {
	// +listType
	Group BuildpackMetadataList `json:"group,omitempty"`
}

// +k8s:openapi-gen=true
type BuildpackMetadata struct {
	Id       string        `json:"id"`
	Version  string        `json:"version"`
	Homepage string        `json:"homepage,omitempty"`
	Order    MetadataOrder `json:"order,omitempty"`
}

func (l MetadataOrder) Include(q BuildpackMetadata) bool {
	for _, bp := range l {
		for _, metadata := range bp.Group {
			if metadata.Id == q.Id && metadata.Version == q.Version {
				return true
			}

			for _, item := range metadata.Order {
				for _, nested := range item.Group {
					if nested.Id == q.Id && nested.Version == q.Version {
						return true
					}
				}
			}
		}
	}

	//
	//for _, bp := range l {
	//	if bp.Id == q.Id && bp.Version == q.Version {
	//		return true
	//	}
	//
	//	for _, group := range bp.Order {
	//		for _, item := range group.Group {
	//			if item.Id == q.Id && item.Version == q.Version {
	//				return true
	//			}
	//		}
	//	}
	//}

	return false
}
