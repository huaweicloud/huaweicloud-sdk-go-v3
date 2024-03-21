package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPremiumHostRequest Request Object
type ListPremiumHostRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *string `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。如果需要一次查全部域名，该参数值填-1。
	Pagesize *string `json:"pagesize,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`

	// 策略名称
	Policyname *string `json:"policyname,omitempty"`

	// 域名防护状态：  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`
}

func (o ListPremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPremiumHostRequest struct{}"
	}

	return strings.Join([]string{"ListPremiumHostRequest", string(data)}, " ")
}
