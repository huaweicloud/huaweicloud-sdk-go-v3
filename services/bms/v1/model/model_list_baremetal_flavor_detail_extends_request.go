package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBaremetalFlavorDetailExtendsRequest struct {

	// 可用区，需要指定可用区（AZ）的名称。请参考地区和终端节点获取。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListBaremetalFlavorDetailExtendsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaremetalFlavorDetailExtendsRequest struct{}"
	}

	return strings.Join([]string{"ListBaremetalFlavorDetailExtendsRequest", string(data)}, " ")
}
