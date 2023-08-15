package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogtankResponse Response Object
type UpdateLogtankResponse struct {

	// 请求的唯一标识。
	RequestId *string `json:"request_id,omitempty"`

	Logtank        *LogtankItem `json:"logtank,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogtankResponse", string(data)}, " ")
}
