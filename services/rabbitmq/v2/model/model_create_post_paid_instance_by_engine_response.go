package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidInstanceByEngineResponse Response Object
type CreatePostPaidInstanceByEngineResponse struct {

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPaidInstanceByEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceByEngineResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceByEngineResponse", string(data)}, " ")
}
