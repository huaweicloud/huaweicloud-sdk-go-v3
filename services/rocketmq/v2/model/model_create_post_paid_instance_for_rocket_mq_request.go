package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidInstanceForRocketMqRequest Request Object
type CreatePostPaidInstanceForRocketMqRequest struct {
	Body *CreateInstanceByEngineReq `json:"body,omitempty"`
}

func (o CreatePostPaidInstanceForRocketMqRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceForRocketMqRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceForRocketMqRequest", string(data)}, " ")
}
