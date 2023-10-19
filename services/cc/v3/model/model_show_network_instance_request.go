package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkInstanceRequest Request Object
type ShowNetworkInstanceRequest struct {

	// 资源的Id。
	Id string `json:"id"`
}

func (o ShowNetworkInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowNetworkInstanceRequest", string(data)}, " ")
}
