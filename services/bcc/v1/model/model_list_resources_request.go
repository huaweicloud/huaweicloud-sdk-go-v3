package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesRequest Request Object
type ListResourcesRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 云服务类型，如ecs,evs,sfsturbo,workspace,rds,gaussdb
	Provider *ListResourcesRequestProvider `json:"provider,omitempty"`

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 标签列表, 标签值的格式为key或key=value,用英文逗号分隔，最多支持5对标签的查询，超过5对的只生效前5对
	Tags *string `json:"tags,omitempty"`

	// CBR存储库ID
	VaultId *string `json:"vault_id,omitempty"`

	// CBR存储库名称
	VaultName *string `json:"vault_name,omitempty"`

	// 资源等级定义ID
	ResourceLevelId *string `json:"resource_level_id,omitempty"`

	// 资源等级名称
	ResourceLevelName *string `json:"resource_level_name,omitempty"`

	// 资源合规结果,complete_match,partial_match,no_match
	ComplianceResult *ListResourcesRequestComplianceResult `json:"compliance_result,omitempty"`

	// 资源保护状态,unprotected,protected
	InventoryResult *ListResourcesRequestInventoryResult `json:"inventory_result,omitempty"`

	// 合规规则类型,local_backup_enabled,local_backup_retention,local_backup_frequency,remote_backup_enabled,remote_backup_retention,local_worm_enabledd,remote_worm_enabled,is_cross_account_backup_forced
	ComplianceRuleType *ListResourcesRequestComplianceRuleType `json:"compliance_rule_type,omitempty"`

	// 合规规则是否符合，必须搭配compliance_rule_type使用
	ComplianceRuleMatched *bool `json:"compliance_rule_matched,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}

type ListResourcesRequestProvider struct {
	value string
}

type ListResourcesRequestProviderEnum struct {
	ECS       ListResourcesRequestProvider
	EVS       ListResourcesRequestProvider
	SFSTURBO  ListResourcesRequestProvider
	WORKSPACE ListResourcesRequestProvider
	RDS       ListResourcesRequestProvider
	GAUSSDB   ListResourcesRequestProvider
}

func GetListResourcesRequestProviderEnum() ListResourcesRequestProviderEnum {
	return ListResourcesRequestProviderEnum{
		ECS: ListResourcesRequestProvider{
			value: "ecs",
		},
		EVS: ListResourcesRequestProvider{
			value: "evs",
		},
		SFSTURBO: ListResourcesRequestProvider{
			value: "sfsturbo",
		},
		WORKSPACE: ListResourcesRequestProvider{
			value: "workspace",
		},
		RDS: ListResourcesRequestProvider{
			value: "rds",
		},
		GAUSSDB: ListResourcesRequestProvider{
			value: "gaussdb",
		},
	}
}

func (c ListResourcesRequestProvider) Value() string {
	return c.value
}

func (c ListResourcesRequestProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesRequestProvider) UnmarshalJSON(b []byte) error {
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

type ListResourcesRequestComplianceResult struct {
	value string
}

type ListResourcesRequestComplianceResultEnum struct {
	COMPLETE_MATCH ListResourcesRequestComplianceResult
	PARTIAL_MATCH  ListResourcesRequestComplianceResult
	NO_MATCH       ListResourcesRequestComplianceResult
}

func GetListResourcesRequestComplianceResultEnum() ListResourcesRequestComplianceResultEnum {
	return ListResourcesRequestComplianceResultEnum{
		COMPLETE_MATCH: ListResourcesRequestComplianceResult{
			value: "complete_match",
		},
		PARTIAL_MATCH: ListResourcesRequestComplianceResult{
			value: "partial_match",
		},
		NO_MATCH: ListResourcesRequestComplianceResult{
			value: "no_match",
		},
	}
}

func (c ListResourcesRequestComplianceResult) Value() string {
	return c.value
}

func (c ListResourcesRequestComplianceResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesRequestComplianceResult) UnmarshalJSON(b []byte) error {
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

type ListResourcesRequestInventoryResult struct {
	value string
}

type ListResourcesRequestInventoryResultEnum struct {
	UNPROTECTED ListResourcesRequestInventoryResult
	PROTECTED   ListResourcesRequestInventoryResult
}

func GetListResourcesRequestInventoryResultEnum() ListResourcesRequestInventoryResultEnum {
	return ListResourcesRequestInventoryResultEnum{
		UNPROTECTED: ListResourcesRequestInventoryResult{
			value: "unprotected",
		},
		PROTECTED: ListResourcesRequestInventoryResult{
			value: "protected",
		},
	}
}

func (c ListResourcesRequestInventoryResult) Value() string {
	return c.value
}

func (c ListResourcesRequestInventoryResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesRequestInventoryResult) UnmarshalJSON(b []byte) error {
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

type ListResourcesRequestComplianceRuleType struct {
	value string
}

type ListResourcesRequestComplianceRuleTypeEnum struct {
	LOCAL_BACKUP_ENABLED           ListResourcesRequestComplianceRuleType
	LOCAL_BACKUP_RETENTION         ListResourcesRequestComplianceRuleType
	LOCAL_BACKUP_FREQUENCY         ListResourcesRequestComplianceRuleType
	REMOTE_BACKUP_ENABLED          ListResourcesRequestComplianceRuleType
	REMOTE_BACKUP_RETENTION        ListResourcesRequestComplianceRuleType
	LOCAL_WORM_ENABLEDD            ListResourcesRequestComplianceRuleType
	REMOTE_WORM_ENABLED            ListResourcesRequestComplianceRuleType
	IS_CROSS_ACCOUNT_BACKUP_FORCED ListResourcesRequestComplianceRuleType
}

func GetListResourcesRequestComplianceRuleTypeEnum() ListResourcesRequestComplianceRuleTypeEnum {
	return ListResourcesRequestComplianceRuleTypeEnum{
		LOCAL_BACKUP_ENABLED: ListResourcesRequestComplianceRuleType{
			value: "local_backup_enabled",
		},
		LOCAL_BACKUP_RETENTION: ListResourcesRequestComplianceRuleType{
			value: "local_backup_retention",
		},
		LOCAL_BACKUP_FREQUENCY: ListResourcesRequestComplianceRuleType{
			value: "local_backup_frequency",
		},
		REMOTE_BACKUP_ENABLED: ListResourcesRequestComplianceRuleType{
			value: "remote_backup_enabled",
		},
		REMOTE_BACKUP_RETENTION: ListResourcesRequestComplianceRuleType{
			value: "remote_backup_retention",
		},
		LOCAL_WORM_ENABLEDD: ListResourcesRequestComplianceRuleType{
			value: "local_worm_enabledd",
		},
		REMOTE_WORM_ENABLED: ListResourcesRequestComplianceRuleType{
			value: "remote_worm_enabled",
		},
		IS_CROSS_ACCOUNT_BACKUP_FORCED: ListResourcesRequestComplianceRuleType{
			value: "is_cross_account_backup_forced",
		},
	}
}

func (c ListResourcesRequestComplianceRuleType) Value() string {
	return c.value
}

func (c ListResourcesRequestComplianceRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesRequestComplianceRuleType) UnmarshalJSON(b []byte) error {
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
