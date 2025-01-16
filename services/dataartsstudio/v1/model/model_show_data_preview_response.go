package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataPreviewResponse Response Object
type ShowDataPreviewResponse struct {

	// 表中数据信息列表
	Rows *[][]interface{} `json:"rows,omitempty"`

	// 字段信息列表
	Schema         *[]interface{} `json:"schema,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDataPreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataPreviewResponse struct{}"
	}

	return strings.Join([]string{"ShowDataPreviewResponse", string(data)}, " ")
}
