package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbNamesResponse Response Object
type ListDbNamesResponse struct {

	// 数据
	Data *[]interface{} `json:"data,omitempty"`

	// 列表大小
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbNamesResponse struct{}"
	}

	return strings.Join([]string{"ListDbNamesResponse", string(data)}, " ")
}
