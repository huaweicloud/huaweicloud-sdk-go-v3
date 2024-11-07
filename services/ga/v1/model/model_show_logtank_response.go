package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogtankResponse Response Object
type ShowLogtankResponse struct {
	Logtank *LogtankDetail `json:"logtank,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankResponse struct{}"
	}

	return strings.Join([]string{"ShowLogtankResponse", string(data)}, " ")
}
