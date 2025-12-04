package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidKafkaInstanceResponse Response Object
type CreatePostPaidKafkaInstanceResponse struct {

	// **参数解释**： 实例ID。 **取值范围**： 不涉及
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPaidKafkaInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidKafkaInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidKafkaInstanceResponse", string(data)}, " ")
}
