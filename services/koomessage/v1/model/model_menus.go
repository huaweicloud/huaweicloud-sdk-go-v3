package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Menus struct {

	// 各子菜单项配置。
	MenuItems []MenuItem `json:"menu_items"`
}

func (o Menus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Menus struct{}"
	}

	return strings.Join([]string{"Menus", string(data)}, " ")
}
