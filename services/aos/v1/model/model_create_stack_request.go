package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStackRequest Request Object
type CreateStackRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	Body *CreateStackRequestBody `json:"body,omitempty"`
}

func (o CreateStackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStackRequest struct{}"
	}

	return strings.Join([]string{"CreateStackRequest", string(data)}, " ")
}
