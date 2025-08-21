package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFileResponse Response Object
type UpdateFileResponse struct {

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 分支名称
	Branch         *string `json:"branch,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFileResponse struct{}"
	}

	return strings.Join([]string{"UpdateFileResponse", string(data)}, " ")
}
