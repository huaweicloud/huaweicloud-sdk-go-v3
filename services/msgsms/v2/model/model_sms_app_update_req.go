package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsAppUpdateReq struct {

	// 应用名称
	AppName string `json:"app_name"`

	// 是否创建测试签名和模板。只有地域为国内时，该字段有效 1. true：是 2. false：否
	CreateSignAndTemplate *bool `json:"create_sign_and_template,omitempty"`

	// 企业项目ID，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称，默认为default
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 地域 1. cn：国内 2. intl：国际
	Region string `json:"region"`

	// 上行回调地址。只有地域为国内时，该字段有效
	UpLinkAddr *string `json:"up_link_addr,omitempty"`
}

func (o SmsAppUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsAppUpdateReq struct{}"
	}

	return strings.Join([]string{"SmsAppUpdateReq", string(data)}, " ")
}
