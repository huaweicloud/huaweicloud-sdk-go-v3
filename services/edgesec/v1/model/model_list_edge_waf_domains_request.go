package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEdgeWafDomainsRequest Request Object
type ListEdgeWafDomainsRequest struct {

	// 页码， 0全查
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页显示的条目数量， waf每批最大查询数量为100
	PageSize *int32 `json:"page_size,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// waf域名配置类别 0 基础信息，1 waf防护配置信息
	Type *ListEdgeWafDomainsRequestType `json:"type,omitempty"`
}

func (o ListEdgeWafDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeWafDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeWafDomainsRequest", string(data)}, " ")
}

type ListEdgeWafDomainsRequestType struct {
	value string
}

type ListEdgeWafDomainsRequestTypeEnum struct {
	E_0 ListEdgeWafDomainsRequestType
	E_1 ListEdgeWafDomainsRequestType
}

func GetListEdgeWafDomainsRequestTypeEnum() ListEdgeWafDomainsRequestTypeEnum {
	return ListEdgeWafDomainsRequestTypeEnum{
		E_0: ListEdgeWafDomainsRequestType{
			value: "0",
		},
		E_1: ListEdgeWafDomainsRequestType{
			value: "1",
		},
	}
}

func (c ListEdgeWafDomainsRequestType) Value() string {
	return c.value
}

func (c ListEdgeWafDomainsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEdgeWafDomainsRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
