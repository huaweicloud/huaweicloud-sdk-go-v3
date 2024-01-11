package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityDataClassificationRuleGroupsRequest Request Object
type ListSecurityDataClassificationRuleGroupsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 规则组名称
	Name *string `json:"name,omitempty"`

	// 规则组创建者
	Creator *string `json:"creator,omitempty"`

	// 排序字段, createdAt, createdBy, updatedAt, updatedBy, name, description
	OrderBy *ListSecurityDataClassificationRuleGroupsRequestOrderBy `json:"order_by,omitempty"`

	// 是否降序
	Desc *bool `json:"desc,omitempty"`
}

func (o ListSecurityDataClassificationRuleGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDataClassificationRuleGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDataClassificationRuleGroupsRequest", string(data)}, " ")
}

type ListSecurityDataClassificationRuleGroupsRequestOrderBy struct {
	value string
}

type ListSecurityDataClassificationRuleGroupsRequestOrderByEnum struct {
	CREATED_AT  ListSecurityDataClassificationRuleGroupsRequestOrderBy
	CREATED_BY  ListSecurityDataClassificationRuleGroupsRequestOrderBy
	UPDATED_AT  ListSecurityDataClassificationRuleGroupsRequestOrderBy
	UPDATED_BY  ListSecurityDataClassificationRuleGroupsRequestOrderBy
	NAME        ListSecurityDataClassificationRuleGroupsRequestOrderBy
	DESCRIPTION ListSecurityDataClassificationRuleGroupsRequestOrderBy
}

func GetListSecurityDataClassificationRuleGroupsRequestOrderByEnum() ListSecurityDataClassificationRuleGroupsRequestOrderByEnum {
	return ListSecurityDataClassificationRuleGroupsRequestOrderByEnum{
		CREATED_AT: ListSecurityDataClassificationRuleGroupsRequestOrderBy{
			value: "createdAt",
		},
		CREATED_BY: ListSecurityDataClassificationRuleGroupsRequestOrderBy{
			value: "createdBy",
		},
		UPDATED_AT: ListSecurityDataClassificationRuleGroupsRequestOrderBy{
			value: "updatedAt",
		},
		UPDATED_BY: ListSecurityDataClassificationRuleGroupsRequestOrderBy{
			value: "updatedBy",
		},
		NAME: ListSecurityDataClassificationRuleGroupsRequestOrderBy{
			value: "name",
		},
		DESCRIPTION: ListSecurityDataClassificationRuleGroupsRequestOrderBy{
			value: "description",
		},
	}
}

func (c ListSecurityDataClassificationRuleGroupsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityDataClassificationRuleGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityDataClassificationRuleGroupsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
