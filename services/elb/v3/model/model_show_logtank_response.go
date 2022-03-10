package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowLogtankResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Logtank        *Logtank `json:"logtank,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankResponse struct{}"
	}

	return strings.Join([]string{"ShowLogtankResponse", string(data)}, " ")
}
