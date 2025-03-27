package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCommandResultResponse Response Object
type UpdateCommandResultResponse struct {

	// 上报服务端命令执行结果成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCommandResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateCommandResultResponse", string(data)}, " ")
}
