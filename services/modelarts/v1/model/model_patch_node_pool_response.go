package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNodePoolResponse Response Object
type PatchNodePoolResponse struct {
	Metadata *PatchNodePoolMetaVo `json:"metadata,omitempty"`

	Spec           *NodePoolSpec `json:"spec,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o PatchNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNodePoolResponse struct{}"
	}

	return strings.Join([]string{"PatchNodePoolResponse", string(data)}, " ")
}
