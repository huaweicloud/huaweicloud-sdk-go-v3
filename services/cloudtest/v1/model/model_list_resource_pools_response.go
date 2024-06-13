package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcePoolsResponse Response Object
type ListResourcePoolsResponse struct {

	// 对外时：success|error; 对内时：ok|failed
	Status *string `json:"status,omitempty"`

	Result *ResultValueListResourcePoolVo `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourcePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcePoolsResponse", string(data)}, " ")
}
