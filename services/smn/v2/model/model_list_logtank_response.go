package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogtankResponse Response Object
type ListLogtankResponse struct {

	// 请求的唯一标识
	RequestId *string `json:"request_id,omitempty"`

	// 云日志信息数量
	Count *int32 `json:"count,omitempty"`

	// 云日志信息列表
	Logtanks       *[]LogtankItem `json:"logtanks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtankResponse struct{}"
	}

	return strings.Join([]string{"ListLogtankResponse", string(data)}, " ")
}
