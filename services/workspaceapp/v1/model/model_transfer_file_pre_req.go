package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferFilePreReq 文件预流转请求体
type TransferFilePreReq struct {

	// 分享者或接收者的用户标识。
	UserName string `json:"user_name"`

	// 待流转的文件名称，不包含完整路径。
	FileName string `json:"file_name"`

	// 文件编码。
	FileCode string `json:"file_code"`

	// 文件提取码。
	FileAccCode string `json:"file_acc_code"`

	// 文件系统名称。
	FileSystemName *string `json:"file_system_name,omitempty"`
}

func (o TransferFilePreReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferFilePreReq struct{}"
	}

	return strings.Join([]string{"TransferFilePreReq", string(data)}, " ")
}
