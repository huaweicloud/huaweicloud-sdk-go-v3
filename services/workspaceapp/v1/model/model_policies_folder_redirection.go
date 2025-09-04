package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesFolderRedirection 配置文件夹重定向。
type PoliciesFolderRedirection struct {

	// 配置文件夹重定向状态： 0: 已禁用 1: 已启用 2: 未配置
	FolderRedirectionStatus *int32 `json:"folder_redirection_status,omitempty"`

	Options *FolderRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesFolderRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFolderRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesFolderRedirection", string(data)}, " ")
}
