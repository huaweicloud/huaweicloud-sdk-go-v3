package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNetworkInstanceRequest Request Object
type DeleteNetworkInstanceRequest struct {

	// 实例ID。
	Id string `json:"id"`
}

func (o DeleteNetworkInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNetworkInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteNetworkInstanceRequest", string(data)}, " ")
}
