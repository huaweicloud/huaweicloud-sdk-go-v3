package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTopicAccessPolicyReq struct {
	// 策略列表。

	Topics []UpdateTopicAccessPolicyTopicsObject `json:"topics"`
}

func (o UpdateTopicAccessPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyReq", string(data)}, " ")
}
