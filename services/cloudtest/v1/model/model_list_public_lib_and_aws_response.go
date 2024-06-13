package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicLibAndAwsResponse Response Object
type ListPublicLibAndAwsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorOfApiTest `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 公共aw信息
	Result *[]GetPublicLibAndAwsResp `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPublicLibAndAwsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicLibAndAwsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicLibAndAwsResponse", string(data)}, " ")
}
