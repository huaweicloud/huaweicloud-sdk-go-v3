package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStackSetRequest Request Object
type CreateStackSetRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	Body *CreateStackSetRequestBody `json:"body,omitempty"`
}

func (o CreateStackSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStackSetRequest struct{}"
	}

	return strings.Join([]string{"CreateStackSetRequest", string(data)}, " ")
}
