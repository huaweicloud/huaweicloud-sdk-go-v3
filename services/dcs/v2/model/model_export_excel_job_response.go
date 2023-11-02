package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportExcelJobResponse Response Object
type ExportExcelJobResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 错误码
	ErrCode *string `json:"err_code,omitempty"`

	// 错误信息
	ErrMsg         *string `json:"err_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportExcelJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportExcelJobResponse struct{}"
	}

	return strings.Join([]string{"ExportExcelJobResponse", string(data)}, " ")
}
