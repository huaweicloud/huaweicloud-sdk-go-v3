package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferFileResponse Response Object
type TransferFileResponse struct {

	// 待流转文件url。
	FileUrl *string `json:"file_url,omitempty"`

	// 文件编码。
	FileCode *string `json:"file_code,omitempty"`

	// 文件提取码。
	FileAccCode *string `json:"file_acc_code,omitempty"`

	// 文件系统名称。
	FileSystemName *string `json:"file_system_name,omitempty"`

	// 文件夹名称。
	FolderName     *string `json:"folder_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TransferFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferFileResponse struct{}"
	}

	return strings.Join([]string{"TransferFileResponse", string(data)}, " ")
}
