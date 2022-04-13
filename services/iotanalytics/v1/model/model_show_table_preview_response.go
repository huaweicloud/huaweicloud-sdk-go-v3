package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTablePreviewResponse struct {
	// 表的列名称和类型。

	Schema *[]interface{} `json:"schema,omitempty"`
	// 预览的表内容。

	Rows           *[]interface{} `json:"rows,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTablePreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTablePreviewResponse struct{}"
	}

	return strings.Join([]string{"ShowTablePreviewResponse", string(data)}, " ")
}
