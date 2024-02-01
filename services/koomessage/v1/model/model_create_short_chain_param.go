package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShortChainParam 生成短链的参数对象。
type CreateShortChainParam struct {

	// 创建解析任务时填写用户唯一标识，手机号码或者任何的唯一标识，唯一标识不超过64个字符。 发送智能信息时则必须填客户的手机号码。样例为：130****0001。
	CustFlag string `json:"cust_flag"`

	// 动态参数。  > 使用动态参数模板时，aim_code_type字段只能为individual。
	DyncParams map[string]string `json:"dync_params,omitempty"`

	// 自定义跳转地址。长度要求不超过256。 > - 未填时，终端用户点击短信原文中的短链后，跳转智能信息模板H5页 > - 已填时，终端用户点击短信原文中的短链后，跳转该字段对应的页面，填写时必须为http或https作为前缀 > - 使用自定义跳转链接功能请联系KooMessage运营人员进行域名备案 > - 自定义短码时即generation_type为2时不支持自定义跳转链接功能，传入的参数值无效
	CustomUrl *string `json:"custom_url,omitempty"`

	// 自定义短码，支持长度为3到10位的数字或大小写字母。样例为：aDC123。 > 自定义短码时即generation_type为2时，此参数为必填。
	CustomShortCode *string `json:"custom_short_code,omitempty"`
}

func (o CreateShortChainParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShortChainParam struct{}"
	}

	return strings.Join([]string{"CreateShortChainParam", string(data)}, " ")
}
