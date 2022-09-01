package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPremiumHostRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *string `json:"page,omitempty" xml:"page"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。
	Pagesize *string `json:"pagesize,omitempty" xml:"pagesize"`

	// 域名
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 策略名称
	Policyname *string `json:"policyname,omitempty" xml:"policyname"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty" xml:"protect_status"`
}

func (o ListPremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPremiumHostRequest struct{}"
	}

	return strings.Join([]string{"ListPremiumHostRequest", string(data)}, " ")
}
