package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用模板
type App struct {
	App *AppDetail `json:"app"`
}

func (o App) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "App struct{}"
	}

	return strings.Join([]string{"App", string(data)}, " ")
}
