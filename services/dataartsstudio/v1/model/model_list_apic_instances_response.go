package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApicInstancesResponse Response Object
type ListApicInstancesResponse struct {

	// 网关实例
	GatewayInstances *[]ApigInstanceDto `json:"gateway_instances,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListApicInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApicInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListApicInstancesResponse", string(data)}, " ")
}
