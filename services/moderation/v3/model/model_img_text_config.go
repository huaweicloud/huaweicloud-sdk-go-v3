package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图文审核场景的黑白词库传入参数设置。
type ImgTextConfig struct {

	// 用户输入的console界面自定义的黑词库名，支持传入多个。
	BlackGlossaryNames *[]string `json:"black_glossary_names,omitempty"`

	// 用户输入的console界面自定义的白词库名，支持传入多个。
	WhiteGlossaryNames *[]string `json:"white_glossary_names,omitempty"`
}

func (o ImgTextConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImgTextConfig struct{}"
	}

	return strings.Join([]string{"ImgTextConfig", string(data)}, " ")
}
