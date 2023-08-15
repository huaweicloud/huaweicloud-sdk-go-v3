package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankResponse Response Object
type CreateLogtankResponse struct {

	// 请求的唯一标识。
	RequestId *string `json:"request_id,omitempty"`

	// 云日志信息唯一的资源标识。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankResponse struct{}"
	}

	return strings.Join([]string{"CreateLogtankResponse", string(data)}, " ")
}
