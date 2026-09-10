package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachShareFilesystemRequestBody struct {

	// 共享文件系统类型，当前仅支持“sfs_turbo”。
	Type string `json:"type"`

	// 共享文件系统id。
	Id string `json:"id"`

	// 云手机服务器id列表。一次最多挂载20台。
	ServerIds []string `json:"server_ids"`

	// 合法的的子目录全路径，不填时默认挂载文件系统的根目录。
	Path *string `json:"path,omitempty"`
}

func (o AttachShareFilesystemRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachShareFilesystemRequestBody struct{}"
	}

	return strings.Join([]string{"AttachShareFilesystemRequestBody", string(data)}, " ")
}
