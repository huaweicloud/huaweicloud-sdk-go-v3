package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackUpMindmapParam struct {

	// 备份名称
	BakName *string `json:"bak_name,omitempty"`

	// 备注
	Memo *string `json:"memo,omitempty"`

	// 脑图id
	MindmapId *string `json:"mindmap_id,omitempty"`
}

func (o BackUpMindmapParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackUpMindmapParam struct{}"
	}

	return strings.Join([]string{"BackUpMindmapParam", string(data)}, " ")
}
