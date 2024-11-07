package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankResponse Response Object
type CreateLogtankResponse struct {
	Logtank *LogtankDetail `json:"logtank,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankResponse struct{}"
	}

	return strings.Join([]string{"CreateLogtankResponse", string(data)}, " ")
}
