package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateProviderRequest Request Object
type CreatePrivateProviderRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	Body *CreatePrivateProviderRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateProviderRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateProviderRequest", string(data)}, " ")
}
