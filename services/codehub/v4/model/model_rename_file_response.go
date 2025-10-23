package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameFileResponse Response Object
type RenameFileResponse struct {

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 分支名称
	Branch         *string `json:"branch,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RenameFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameFileResponse struct{}"
	}

	return strings.Join([]string{"RenameFileResponse", string(data)}, " ")
}
