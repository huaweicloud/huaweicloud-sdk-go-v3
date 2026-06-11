package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesFolderRedirectionV2 配置文件夹重定向(v2)。
type PoliciesFolderRedirectionV2 struct {

	// 配置文件夹重定向状态： 0: 关闭 1: 已启用
	FolderRedirectionV2Status *int32 `json:"folder_redirection_v2_status,omitempty"`

	Options *FolderRedirectionV2Options `json:"options,omitempty"`
}

func (o PoliciesFolderRedirectionV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFolderRedirectionV2 struct{}"
	}

	return strings.Join([]string{"PoliciesFolderRedirectionV2", string(data)}, " ")
}
