package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IDc struct {
	Name string `json:"name"`

	IrackNum int32 `json:"irack_num"`

	Id string `json:"id"`

	Region *string `json:"region,omitempty"`

	Deccription *string `json:"deccription,omitempty"`
}

func (o IDc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IDc struct{}"
	}

	return strings.Join([]string{"IDc", string(data)}, " ")
}
