package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Subnet struct {

	// 子网的网络ID。
	SubnetId string `json:"subnet_id"`
}

func (o Subnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subnet struct{}"
	}

	return strings.Join([]string{"Subnet", string(data)}, " ")
}
