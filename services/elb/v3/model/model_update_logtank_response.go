package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogtankResponse Response Object
type UpdateLogtankResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Logtank        *Logtank `json:"logtank,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogtankResponse", string(data)}, " ")
}
