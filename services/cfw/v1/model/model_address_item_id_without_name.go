package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressItemIdWithoutName struct {

	// 地址组成员id
	Id *string `json:"id,omitempty"`
}

func (o AddressItemIdWithoutName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemIdWithoutName struct{}"
	}

	return strings.Join([]string{"AddressItemIdWithoutName", string(data)}, " ")
}
