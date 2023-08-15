package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MenusMode struct {

	// 各子菜单项配置。
	MenuItems []MenuItem `json:"menu_items"`
}

func (o MenusMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MenusMode struct{}"
	}

	return strings.Join([]string{"MenusMode", string(data)}, " ")
}
