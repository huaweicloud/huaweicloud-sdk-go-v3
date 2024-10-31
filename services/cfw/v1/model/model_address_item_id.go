package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressItemId struct {

	// 地址组成员id
	Id *string `json:"id,omitempty"`

	// 地址组成员名称
	Name *string `json:"name,omitempty"`
}

func (o AddressItemId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemId struct{}"
	}

	return strings.Join([]string{"AddressItemId", string(data)}, " ")
}
