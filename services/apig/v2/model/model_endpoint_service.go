package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointService struct {

	// 终端节点服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o EndpointService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointService struct{}"
	}

	return strings.Join([]string{"EndpointService", string(data)}, " ")
}
