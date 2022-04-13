package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建网络实例的请求体。
type CreateNetworkInstanceRequestBody struct {
	NetworkInstance *CreateNetworkInstance `json:"network_instance"`
}

func (o CreateNetworkInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNetworkInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNetworkInstanceRequestBody", string(data)}, " ")
}
