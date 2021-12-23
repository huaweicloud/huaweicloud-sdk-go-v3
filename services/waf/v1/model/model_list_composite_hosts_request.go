package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCompositeHostsRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 分页查询时，返回第几页数据。范围0-100000，默认值为1，表示返回第1页数据。

	Page *int32 `json:"page,omitempty"`
	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 域名名称

	Hostname *string `json:"hostname,omitempty"`
	// 防护策略名称

	Policyname *string `json:"policyname,omitempty"`
	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测

	ProtectStatus *ListCompositeHostsRequestProtectStatus `json:"protect_status,omitempty"`
	// 域名所属WAF模式

	WafType *string `json:"waf_type,omitempty"`
	// 域名是否使用HTTPS

	IsHttps *bool `json:"is_https,omitempty"`
}

func (o ListCompositeHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompositeHostsRequest struct{}"
	}

	return strings.Join([]string{"ListCompositeHostsRequest", string(data)}, " ")
}

type ListCompositeHostsRequestProtectStatus struct {
	value int32
}

type ListCompositeHostsRequestProtectStatusEnum struct {
	E_MINUS_1 ListCompositeHostsRequestProtectStatus
	E_0       ListCompositeHostsRequestProtectStatus
	E_1       ListCompositeHostsRequestProtectStatus
}

func GetListCompositeHostsRequestProtectStatusEnum() ListCompositeHostsRequestProtectStatusEnum {
	return ListCompositeHostsRequestProtectStatusEnum{
		E_MINUS_1: ListCompositeHostsRequestProtectStatus{
			value: -1,
		}, E_0: ListCompositeHostsRequestProtectStatus{
			value: 0,
		}, E_1: ListCompositeHostsRequestProtectStatus{
			value: 1,
		},
	}
}

func (c ListCompositeHostsRequestProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCompositeHostsRequestProtectStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
