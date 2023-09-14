package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupplementdataResponse Response Object
type ListSupplementdataResponse struct {
	Msg *string `json:"msg,omitempty"`

	// 包含若干补数据实例信息
	Rows *[]SupplementDataResp `json:"rows,omitempty"`

	// 查询是否成功，取值为true或者false
	Success *bool `json:"success,omitempty"`

	// 补数据实例数量
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSupplementdataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupplementdataResponse struct{}"
	}

	return strings.Join([]string{"ListSupplementdataResponse", string(data)}, " ")
}
