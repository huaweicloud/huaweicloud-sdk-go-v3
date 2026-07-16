package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNodePoolMetaVo 节点池的metadata信息。
type PatchNodePoolMetaVo struct {
	Annotations *PatchNodePoolAnnotations `json:"annotations,omitempty"`
}

func (o PatchNodePoolMetaVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNodePoolMetaVo struct{}"
	}

	return strings.Join([]string{"PatchNodePoolMetaVo", string(data)}, " ")
}
