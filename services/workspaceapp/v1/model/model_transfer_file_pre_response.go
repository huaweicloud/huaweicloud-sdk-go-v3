package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferFilePreResponse Response Object
type TransferFilePreResponse struct {

	// 待流转文件url。
	TargetFileUrl *string `json:"target_file_url,omitempty"`

	// 文件系统名称。
	FileSystemName *string `json:"file_system_name,omitempty"`

	// 用户可自定义的文件夹名称。
	FolderName     *string `json:"folder_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TransferFilePreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferFilePreResponse struct{}"
	}

	return strings.Join([]string{"TransferFilePreResponse", string(data)}, " ")
}
