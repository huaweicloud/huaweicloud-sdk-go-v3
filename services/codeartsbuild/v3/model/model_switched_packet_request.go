package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchedPacketRequest Request Object
type SwitchedPacketRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 源分组编号
	SourceGroupId string `json:"source_group_id"`

	// 目标分组编号
	TargetGroupId string `json:"target_group_id"`
}

func (o SwitchedPacketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchedPacketRequest struct{}"
	}

	return strings.Join([]string{"SwitchedPacketRequest", string(data)}, " ")
}
