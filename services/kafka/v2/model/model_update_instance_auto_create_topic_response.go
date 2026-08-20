package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceAutoCreateTopicResponse Response Object
type UpdateInstanceAutoCreateTopicResponse struct {

	// **参数解释**： 是否开启自动创建Topic功能。 **取值范围**： - true：开启自动创建Topic功能。 - false：关闭自动创建Topic功能。
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateInstanceAutoCreateTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAutoCreateTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAutoCreateTopicResponse", string(data)}, " ")
}
