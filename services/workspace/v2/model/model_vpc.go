package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Vpc vpc信息。
type Vpc struct {

	// 桌面所在vpcId。
	Id *string `json:"id,omitempty"`

	// 桌面所在vpc名称。
	Name *string `json:"name,omitempty"`

	// 桌面所在vpc网段。
	Cidr *string `json:"cidr,omitempty"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}
