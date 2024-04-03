package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountGlobalEipsResponse Response Object
type CountGlobalEipsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	GlobalEip *CountGlobalEips `json:"global_eip,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountGlobalEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGlobalEipsResponse struct{}"
	}

	return strings.Join([]string{"CountGlobalEipsResponse", string(data)}, " ")
}
