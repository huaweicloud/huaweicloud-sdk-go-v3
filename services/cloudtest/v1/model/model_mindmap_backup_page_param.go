package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MindmapBackupPageParam struct {

	// 备份名称
	BakName *string `json:"bak_name,omitempty"`

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	// 脑图ID
	MindmapId *string `json:"mindmap_id,omitempty"`

	// 备份类型
	Type *string `json:"type,omitempty"`
}

func (o MindmapBackupPageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindmapBackupPageParam struct{}"
	}

	return strings.Join([]string{"MindmapBackupPageParam", string(data)}, " ")
}
