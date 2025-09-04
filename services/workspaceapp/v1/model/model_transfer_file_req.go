package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferFileReq 文件分享
type TransferFileReq struct {

	// 操作方式, 0:个人->共享, 1:共享->个人。
	TransferType int32 `json:"transfer_type"`

	// 分享者或接收者的用户标识。
	UserName string `json:"user_name"`

	// 待流转文件url。
	FileUrl string `json:"file_url"`

	// 文件编码(仅接收文件使用，从分享返回体获取)。
	FileCode *string `json:"file_code,omitempty"`

	// 文件提取码(仅接收文件使用，从分享返回体获取)。
	FileAccCode *string `json:"file_acc_code,omitempty"`

	// 目标文件url。
	TargetFileUrl *string `json:"target_file_url,omitempty"`

	// 文件系统名称。
	FileSystemName *string `json:"file_system_name,omitempty"`
}

func (o TransferFileReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferFileReq struct{}"
	}

	return strings.Join([]string{"TransferFileReq", string(data)}, " ")
}
