package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipResponse Response Object
type BatchCreateGlobalEipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 响应对象
	GlobalEips *[]BatchCreateGlobalEipJob `json:"global_eips,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateGlobalEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipResponse", string(data)}, " ")
}
