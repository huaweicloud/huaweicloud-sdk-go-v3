package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新网络实例的请求体。
type UpdateNetworkInstanceRequestBody struct {
	NetworkInstance *UpdateNetworkInstance `json:"network_instance" xml:"network_instance"`
}

func (o UpdateNetworkInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNetworkInstanceRequestBody", string(data)}, " ")
}
