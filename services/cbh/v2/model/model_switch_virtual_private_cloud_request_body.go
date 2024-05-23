package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchVirtualPrivateCloudRequestBody struct {

	// 需要切换vpc的实例id
	ServerId string `json:"server_id"`

	Network *NetworkInfoCreate `json:"network"`
}

func (o SwitchVirtualPrivateCloudRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchVirtualPrivateCloudRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchVirtualPrivateCloudRequestBody", string(data)}, " ")
}
