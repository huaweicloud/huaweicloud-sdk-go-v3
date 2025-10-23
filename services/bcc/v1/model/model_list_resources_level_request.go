package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesLevelRequest Request Object
type ListResourcesLevelRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 云服务类型：ecs-云服务器，evs-云硬盘，sfsturbo-高性能文件存储，workspace-云桌面，rds-关系型数据库，gaussdb-高斯数据库
	Provider *ListResourcesLevelRequestProvider `json:"provider,omitempty"`

	// 合规规则ID
	ComplianceRuleId *string `json:"compliance_rule_id,omitempty"`

	// 本次请求返回的最大记录条数。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListResourcesLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesLevelRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesLevelRequest", string(data)}, " ")
}

type ListResourcesLevelRequestProvider struct {
	value string
}

type ListResourcesLevelRequestProviderEnum struct {
	ECS       ListResourcesLevelRequestProvider
	EVS       ListResourcesLevelRequestProvider
	SFSTURBO  ListResourcesLevelRequestProvider
	WORKSPACE ListResourcesLevelRequestProvider
	RDS       ListResourcesLevelRequestProvider
	GAUSSDB   ListResourcesLevelRequestProvider
}

func GetListResourcesLevelRequestProviderEnum() ListResourcesLevelRequestProviderEnum {
	return ListResourcesLevelRequestProviderEnum{
		ECS: ListResourcesLevelRequestProvider{
			value: "ecs",
		},
		EVS: ListResourcesLevelRequestProvider{
			value: "evs",
		},
		SFSTURBO: ListResourcesLevelRequestProvider{
			value: "sfsturbo",
		},
		WORKSPACE: ListResourcesLevelRequestProvider{
			value: "workspace",
		},
		RDS: ListResourcesLevelRequestProvider{
			value: "rds",
		},
		GAUSSDB: ListResourcesLevelRequestProvider{
			value: "gaussdb",
		},
	}
}

func (c ListResourcesLevelRequestProvider) Value() string {
	return c.value
}

func (c ListResourcesLevelRequestProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesLevelRequestProvider) UnmarshalJSON(b []byte) error {
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
