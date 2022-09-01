package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTablePreviewResponse struct {

	// 表的列名称和类型。
	Schema *[]interface{} `json:"schema,omitempty" xml:"schema"`

	// 预览的表内容。
	Rows           *[]interface{} `json:"rows,omitempty" xml:"rows"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTablePreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTablePreviewResponse struct{}"
	}

	return strings.Join([]string{"ShowTablePreviewResponse", string(data)}, " ")
}
