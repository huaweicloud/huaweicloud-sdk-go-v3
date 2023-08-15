package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImgTextConfig 图文审核场景的黑白词库传入参数设置。
type ImgTextConfig struct {

	// 检测时使用的自定义黑名单词库列表。
	BlackGlossaryNames *[]string `json:"black_glossary_names,omitempty"`

	// 检测时使用的自定义白名单词库列表。
	WhiteGlossaryNames *[]string `json:"white_glossary_names,omitempty"`
}

func (o ImgTextConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImgTextConfig struct{}"
	}

	return strings.Join([]string{"ImgTextConfig", string(data)}, " ")
}
