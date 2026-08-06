package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNameListResponse Response Object
type ShowNameListResponse struct {

	// 数据
	Data *[]interface{} `json:"data,omitempty"`

	// 列表大小
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowNameListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNameListResponse struct{}"
	}

	return strings.Join([]string{"ShowNameListResponse", string(data)}, " ")
}
