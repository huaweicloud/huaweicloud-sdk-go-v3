package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleTagDto struct {

	// 标签名称
	Name *string `json:"name,omitempty"`

	// 标签描述
	Message *string `json:"message,omitempty"`

	// 目标commit_id
	Target *string `json:"target,omitempty"`

	GpgKey *GpgKeyDto `json:"gpgKey,omitempty"`

	// 是否可以被删除
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否可以被查看
	CanRead *bool `json:"can_read,omitempty"`

	// 是否可以被下载
	CanDownload *bool `json:"can_download,omitempty"`
}

func (o SimpleTagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleTagDto struct{}"
	}

	return strings.Join([]string{"SimpleTagDto", string(data)}, " ")
}
