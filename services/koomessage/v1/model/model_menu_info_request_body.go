package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MenuInfoRequestBody struct {
	Menu *MenusMode `json:"menu"`
}

func (o MenuInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MenuInfoRequestBody struct{}"
	}

	return strings.Join([]string{"MenuInfoRequestBody", string(data)}, " ")
}
