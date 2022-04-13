package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建负载均衡器的请求体
type CreateLoadbalancerReq struct {
	// 负载均衡器所在的项目ID。

	TenantId *string `json:"tenant_id,omitempty"`
	// 负载均衡器名称。

	Name *string `json:"name,omitempty"`
	// 负载均衡器的描述信息

	Description *string `json:"description,omitempty"`
	// 负载均衡器所在的子网ID

	VipSubnetId string `json:"vip_subnet_id"`
	// 负载均衡器的虚拟IP。

	VipAddress *string `json:"vip_address,omitempty"`
	// 负载均衡器的供应者名称。只支持vlb

	Provider *CreateLoadbalancerReqProvider `json:"provider,omitempty"`
	// 负载均衡器的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 企业项目ID。创建负载均衡器时，给负载均衡器绑定企业项目ID。取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参见《企业管理用户指南》。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateLoadbalancerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerReq struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerReq", string(data)}, " ")
}

type CreateLoadbalancerReqProvider struct {
	value string
}

type CreateLoadbalancerReqProviderEnum struct {
	VLB CreateLoadbalancerReqProvider
}

func GetCreateLoadbalancerReqProviderEnum() CreateLoadbalancerReqProviderEnum {
	return CreateLoadbalancerReqProviderEnum{
		VLB: CreateLoadbalancerReqProvider{
			value: "vlb",
		},
	}
}

func (c CreateLoadbalancerReqProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadbalancerReqProvider) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
