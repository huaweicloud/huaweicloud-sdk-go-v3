package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogtankResponse Response Object
type UpdateLogtankResponse struct {
	Logtank *LogtankDetail `json:"logtank,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogtankResponse", string(data)}, " ")
}
