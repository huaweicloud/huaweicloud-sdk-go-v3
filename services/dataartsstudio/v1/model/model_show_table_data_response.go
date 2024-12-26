package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableDataResponse Response Object
type ShowTableDataResponse struct {

	// 表中数据信息列表
	Rows *[][]interface{} `json:"rows,omitempty"`

	// 字段信息列表
	Schema         *[]interface{} `json:"schema,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTableDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableDataResponse struct{}"
	}

	return strings.Join([]string{"ShowTableDataResponse", string(data)}, " ")
}
