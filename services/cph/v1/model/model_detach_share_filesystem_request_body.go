package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DetachShareFilesystemRequestBody struct {

	// 共享文件系统类型，当前仅支持“sfs_turbo”。
	Type string `json:"type"`

	// 共享文件系统id。
	Id string `json:"id"`

	// 云手机服务器id列表。一次最多卸载20台。
	ServerIds []string `json:"server_ids"`
}

func (o DetachShareFilesystemRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachShareFilesystemRequestBody struct{}"
	}

	return strings.Join([]string{"DetachShareFilesystemRequestBody", string(data)}, " ")
}
