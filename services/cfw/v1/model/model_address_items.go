package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressItems struct {

	// 地址组成员id列表
	Items *[]IdObject `json:"items,omitempty"`
}

func (o AddressItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItems struct{}"
	}

	return strings.Join([]string{"AddressItems", string(data)}, " ")
}
