package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestTypesResponse Response Object
type ListTestTypesResponse struct {

	// 对外时：success|error; 对内时：ok|failed
	Status *string `json:"status,omitempty"`

	Result *ResultValueListIntegerIdAndNameVo `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId *string `json:"request_id,omitempty"`

	// 对内接口才有此属性
	ServerAddress  *string `json:"server_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTestTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestTypesResponse struct{}"
	}

	return strings.Join([]string{"ListTestTypesResponse", string(data)}, " ")
}
