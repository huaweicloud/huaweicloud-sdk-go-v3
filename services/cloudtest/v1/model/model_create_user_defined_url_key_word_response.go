package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserDefinedUrlKeyWordResponse Response Object
type CreateUserDefinedUrlKeyWordResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorObject `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 结果
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUserDefinedUrlKeyWordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserDefinedUrlKeyWordResponse struct{}"
	}

	return strings.Join([]string{"CreateUserDefinedUrlKeyWordResponse", string(data)}, " ")
}
