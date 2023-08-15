package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndividualContentParam 智能信息基础版参数名称。
type IndividualContentParam struct {

	// 智能信息基础版参数名称。
	ParamName string `json:"param_name"`

	// 智能信息基础版参数类型。txt：纯文字动参。
	ContentType string `json:"content_type"`

	// 智能信息基础版参数源。txt：内容源自纯文字。
	ContentSource string `json:"content_source"`

	// 智能信息基础版参数内容，填写经过utf-8编码的文字。
	ContentDetail string `json:"content_detail"`
}

func (o IndividualContentParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndividualContentParam struct{}"
	}

	return strings.Join([]string{"IndividualContentParam", string(data)}, " ")
}
