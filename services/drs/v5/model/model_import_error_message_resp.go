package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportErrorMessageResp 导入失败详情响应体。
type ImportErrorMessageResp struct {

	// sheet名称。
	SheetName *string `json:"sheet_name,omitempty"`

	// 行号。
	RowRum *string `json:"row_rum,omitempty"`

	// 字段值。
	Value *string `json:"value,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ImportErrorMessageResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportErrorMessageResp struct{}"
	}

	return strings.Join([]string{"ImportErrorMessageResp", string(data)}, " ")
}
