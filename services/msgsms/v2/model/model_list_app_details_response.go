package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppDetailsResponse struct {

	// 查询结果
	Results *[]SmsAppQueryResp `json:"results,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListAppDetailsResponse", string(data)}, " ")
}
