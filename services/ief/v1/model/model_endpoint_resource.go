package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointResource 消息端点资源属性
type EndpointResource struct {

	// 消息端点资源。 示例：- dis: {\"channel\": \"dis channel name\"} - servicebus: {\"path\": \"/request path\"} - apigw: {\"resource\": \"http://ssss.com\"} - eventbus: {\"topic\": \"/xxxx\"}
	Resource *string `json:"resource,omitempty"`
}

func (o EndpointResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointResource struct{}"
	}

	return strings.Join([]string{"EndpointResource", string(data)}, " ")
}
