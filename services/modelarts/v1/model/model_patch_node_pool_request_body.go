package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNodePoolRequestBody 更新节点池的请求体。
type PatchNodePoolRequestBody struct {
	Metadata *PatchNodePoolMetaVo `json:"metadata,omitempty"`

	Spec *NodePoolSpec `json:"spec"`
}

func (o PatchNodePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNodePoolRequestBody struct{}"
	}

	return strings.Join([]string{"PatchNodePoolRequestBody", string(data)}, " ")
}
