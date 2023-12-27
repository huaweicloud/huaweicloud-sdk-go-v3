package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LanguageCheckInfoReq struct {

	// 目标语言
	TargetLanguage string `json:"target_language"`

	// 用户传来的剧本文本信息
	ShootScript []LiveShootScriptItem `json:"shoot_script"`
}

func (o LanguageCheckInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LanguageCheckInfoReq struct{}"
	}

	return strings.Join([]string{"LanguageCheckInfoReq", string(data)}, " ")
}
