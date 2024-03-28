package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunIefJobActionCallBackRequestBody IEF Flink job action回调的请求body体。
type RunIefJobActionCallBackRequestBody struct {

	// 消息id
	MessageId string `json:"message_id"`

	State *State `json:"state,omitempty"`
}

func (o RunIefJobActionCallBackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunIefJobActionCallBackRequestBody struct{}"
	}

	return strings.Join([]string{"RunIefJobActionCallBackRequestBody", string(data)}, " ")
}
