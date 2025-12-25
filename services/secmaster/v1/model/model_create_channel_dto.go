package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateChannelDto struct {
	Action *ChannelAction `json:"action,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// UUID
	GroupId string `json:"group_id"`

	// 相关描述信息
	Input []CreateModuleVo `json:"input"`

	// 相关描述信息
	Nodes []NodeVo `json:"nodes"`

	// 相关描述信息
	Output []CreateModuleVo `json:"output"`

	// UUID
	ParserId string `json:"parser_id"`

	// 相关标题
	Title string `json:"title"`
}

func (o CreateChannelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChannelDto struct{}"
	}

	return strings.Join([]string{"CreateChannelDto", string(data)}, " ")
}
