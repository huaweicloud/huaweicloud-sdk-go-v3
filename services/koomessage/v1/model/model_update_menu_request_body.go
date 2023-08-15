package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMenuRequestBody struct {
	Menu *Menus `json:"menu"`

	// 修改原因。
	ChangeReason string `json:"change_reason"`
}

func (o UpdateMenuRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMenuRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMenuRequestBody", string(data)}, " ")
}
