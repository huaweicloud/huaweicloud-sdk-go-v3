package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAddressSetResponseData struct {

	// 地址组id
	Id *string `json:"id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`
}

func (o UpdateAddressSetResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetResponseData struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetResponseData", string(data)}, " ")
}
