package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShareFolderReq 创建共享存储声明
type CreateShareFolderReq struct {

	// - 仅支持创建单层级的文件夹 - 单个文件夹名称仅支持以下字符: 英文字母、数字、空格、下划线、中划线 - 名称不能超过32字符 - 不能为全空格或者以空格开头
	FolderName *string `json:"folder_name,omitempty"`
}

func (o CreateShareFolderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareFolderReq struct{}"
	}

	return strings.Join([]string{"CreateShareFolderReq", string(data)}, " ")
}
