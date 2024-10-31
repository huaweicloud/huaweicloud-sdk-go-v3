package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BlackWhiteListId struct {

	// 黑白名单id
	Id *string `json:"id,omitempty"`

	// 黑白名单名称，为黑白名单的地址
	Name *string `json:"name,omitempty"`
}

func (o BlackWhiteListId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteListId struct{}"
	}

	return strings.Join([]string{"BlackWhiteListId", string(data)}, " ")
}
