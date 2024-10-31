package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressSetId struct {

	// 地址组id
	Id *string `json:"id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`
}

func (o AddressSetId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressSetId struct{}"
	}

	return strings.Join([]string{"AddressSetId", string(data)}, " ")
}
