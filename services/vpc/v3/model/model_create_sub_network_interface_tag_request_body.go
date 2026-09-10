package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubNetworkInterfaceTagRequestBody This is a auto create Body Object
type CreateSubNetworkInterfaceTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSubNetworkInterfaceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceTagRequestBody", string(data)}, " ")
}
