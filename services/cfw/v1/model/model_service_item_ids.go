package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceItemIds struct {

	// 服务组成员id列表
	Items *[]IdObject `json:"items,omitempty"`
}

func (o ServiceItemIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItemIds struct{}"
	}

	return strings.Join([]string{"ServiceItemIds", string(data)}, " ")
}
