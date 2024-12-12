package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateModuleRequest Request Object
type CreatePrivateModuleRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	Body *CreatePrivateModuleRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateModuleRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateModuleRequest", string(data)}, " ")
}
