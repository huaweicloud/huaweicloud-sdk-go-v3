package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourceUnderNodeResponse struct {

	// 分页查询的数据。
	Data *[]interface{} `json:"data,omitempty"`

	// 分页信息。
	PageInfo       *interface{} `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListResourceUnderNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceUnderNodeResponse struct{}"
	}

	return strings.Join([]string{"ListResourceUnderNodeResponse", string(data)}, " ")
}
