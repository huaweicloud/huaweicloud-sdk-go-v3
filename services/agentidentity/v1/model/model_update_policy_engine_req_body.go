package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyEngineReqBody struct {

	// 策略集的更新描述。
	Description *string `json:"description,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdatePolicyEngineReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyEngineReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyEngineReqBody", string(data)}, " ")
}
