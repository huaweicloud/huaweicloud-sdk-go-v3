package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteShareFilesRequestBody struct {

	// 云手机服务器ID列表。
	ServerIds []string `json:"server_ids"`

	// 所需删除的共享存储文件绝对路径。以/开头，最大长度4096字节，目前只支持大小写字母、数字、点（.）、斜线（/）、中划线（-）、空格、下划线（_）、等号（=），不支持中文。路径中不能包含.. 上层目录路径，防止跨目录攻击。仅删除共享存储接口使用。
	FilePaths []string `json:"file_paths"`
}

func (o DeleteShareFilesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareFilesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteShareFilesRequestBody", string(data)}, " ")
}
