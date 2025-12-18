package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEndpointServiceServerResourceRequestBody 终端节点服务添加后端服务资源请求结构体。
type AddEndpointServiceServerResourceRequestBody struct {

	// 后端elb实例ID和elb所在的az信息列表。
	ServerResources []ServerResource `json:"server_resources"`
}

func (o AddEndpointServiceServerResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEndpointServiceServerResourceRequestBody struct{}"
	}

	return strings.Join([]string{"AddEndpointServiceServerResourceRequestBody", string(data)}, " ")
}
