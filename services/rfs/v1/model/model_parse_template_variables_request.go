package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseTemplateVariablesRequest Request Object
type ParseTemplateVariablesRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	Body *ParseTemplateVariablesRequestBody `json:"body,omitempty"`
}

func (o ParseTemplateVariablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseTemplateVariablesRequest struct{}"
	}

	return strings.Join([]string{"ParseTemplateVariablesRequest", string(data)}, " ")
}
