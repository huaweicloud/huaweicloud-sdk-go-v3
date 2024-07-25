package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityDynamicMaskingPoliciesRequest Request Object
type ListSecurityDynamicMaskingPoliciesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 动态脱敏策略名称。
	Name *string `json:"name,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 排序参数，UPDATE_TIME。
	OrderBy *ListSecurityDynamicMaskingPoliciesRequestOrderBy `json:"order_by,omitempty"`

	// 是否升序（仅指定排序参数时有效）。
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityDynamicMaskingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDynamicMaskingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDynamicMaskingPoliciesRequest", string(data)}, " ")
}

type ListSecurityDynamicMaskingPoliciesRequestOrderBy struct {
	value string
}

type ListSecurityDynamicMaskingPoliciesRequestOrderByEnum struct {
	UPDATE_TIME ListSecurityDynamicMaskingPoliciesRequestOrderBy
}

func GetListSecurityDynamicMaskingPoliciesRequestOrderByEnum() ListSecurityDynamicMaskingPoliciesRequestOrderByEnum {
	return ListSecurityDynamicMaskingPoliciesRequestOrderByEnum{
		UPDATE_TIME: ListSecurityDynamicMaskingPoliciesRequestOrderBy{
			value: "UPDATE_TIME",
		},
	}
}

func (c ListSecurityDynamicMaskingPoliciesRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityDynamicMaskingPoliciesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityDynamicMaskingPoliciesRequestOrderBy) UnmarshalJSON(b []byte) error {
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
