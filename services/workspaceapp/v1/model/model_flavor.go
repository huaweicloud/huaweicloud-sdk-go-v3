package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Flavor struct {

	// 服务器对应的规格ID
	Id *string `json:"id,omitempty"`

	// 服务器对应规格的相关标记快捷链接信息
	Links *[]FlavorLink `json:"links,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
