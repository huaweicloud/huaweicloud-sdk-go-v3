package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSignatureDetailsResponse Response Object
type ListSignatureDetailsResponse struct {

	// 查询结果
	Results *[]SmsSignatureResp `json:"results,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSignatureDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSignatureDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSignatureDetailsResponse", string(data)}, " ")
}
