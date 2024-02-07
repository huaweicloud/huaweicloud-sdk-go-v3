package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalEipResponse Response Object
type ShowGlobalEipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	GlobalEip *ShowGlobalEip `json:"global_eip,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGlobalEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipResponse", string(data)}, " ")
}
