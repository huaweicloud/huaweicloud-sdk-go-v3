package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLabelPageReq 标签页面请求体
type CreateLabelPageReq struct {

	// 标签页面标题，正则匹配中文，英文字母和数字及下划线
	Name string `json:"name"`

	Feature *FeatureEnum `json:"feature"`

	// 标签页面包含的标签值，正则匹配中文，英文字母和数字及下划线
	Labels []string `json:"labels"`
}

func (o CreateLabelPageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelPageReq struct{}"
	}

	return strings.Join([]string{"CreateLabelPageReq", string(data)}, " ")
}
