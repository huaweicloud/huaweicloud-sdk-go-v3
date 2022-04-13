package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateLogtankResponse struct {
	Logtank *Logtank `json:"logtank,omitempty"`
	// 请求ID。  注：自动生成 。

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
