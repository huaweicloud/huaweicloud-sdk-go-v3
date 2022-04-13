package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteNetworkInstanceRequest struct {
	// 网络实例ID。

	Id string `json:"id"`
}

func (o DeleteNetworkInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNetworkInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteNetworkInstanceRequest", string(data)}, " ")
}
