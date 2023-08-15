package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInfoRef Response of Playbook Info Ref
type PlaybookInfoRef struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// Id value
	VersionId *string `json:"version_id,omitempty"`

	// Id value
	Name *string `json:"name,omitempty"`

	// version
	Version *string `json:"version,omitempty"`
}

func (o PlaybookInfoRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInfoRef struct{}"
	}

	return strings.Join([]string{"PlaybookInfoRef", string(data)}, " ")
}
