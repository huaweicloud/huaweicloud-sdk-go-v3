package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateLogtankResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Logtank        *Logtank `json:"logtank,omitempty" xml:"logtank"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogtankResponse", string(data)}, " ")
}
