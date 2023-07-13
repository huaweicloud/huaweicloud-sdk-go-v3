package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagSatisfactionRequest Request Object
type TagSatisfactionRequest struct {

	// qabot编号，UUID格式。
	QabotId string `json:"qabot_id"`

	// 请求ID，由问答机器人会话生成。
	RequestId string `json:"request_id"`

	Body *PostSatisfactionReq `json:"body,omitempty"`
}

func (o TagSatisfactionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagSatisfactionRequest struct{}"
	}

	return strings.Join([]string{"TagSatisfactionRequest", string(data)}, " ")
}
