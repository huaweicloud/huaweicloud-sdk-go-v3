package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCaseResultResponse Response Object
type ShowCaseResultResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueListTestResultVo `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId *string `json:"request_id,omitempty"`

	// 对内接口才有此属性
	ServerAddress  *string `json:"server_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCaseResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCaseResultResponse struct{}"
	}

	return strings.Join([]string{"ShowCaseResultResponse", string(data)}, " ")
}
