package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityDataClassificationRulesRequest Request Object
type ListSecurityDataClassificationRulesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 查询偏移
	Offset *int32 `json:"offset,omitempty"`

	// 查询一页限制
	Limit *int32 `json:"limit,omitempty"`

	// 密级
	SecrecyLevel *string `json:"secrecy_level,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则创建者
	Creator *string `json:"creator,omitempty"`

	// 规则是否开启
	Enable *bool `json:"enable,omitempty"`

	// 开始日期
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束日期
	EndTime *int64 `json:"end_time,omitempty"`

	// 排序字段, createdAt, createdBy, updatedAt, updatedBy, name, description
	OrderBy *ListSecurityDataClassificationRulesRequestOrderBy `json:"order_by,omitempty"`

	// 排序规则
	Desc *bool `json:"desc,omitempty"`
}

func (o ListSecurityDataClassificationRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDataClassificationRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDataClassificationRulesRequest", string(data)}, " ")
}

type ListSecurityDataClassificationRulesRequestOrderBy struct {
	value string
}

type ListSecurityDataClassificationRulesRequestOrderByEnum struct {
	CREATED_AT  ListSecurityDataClassificationRulesRequestOrderBy
	CREATED_BY  ListSecurityDataClassificationRulesRequestOrderBy
	UPDATED_AT  ListSecurityDataClassificationRulesRequestOrderBy
	UPDATED_BY  ListSecurityDataClassificationRulesRequestOrderBy
	NAME        ListSecurityDataClassificationRulesRequestOrderBy
	DESCRIPTION ListSecurityDataClassificationRulesRequestOrderBy
}

func GetListSecurityDataClassificationRulesRequestOrderByEnum() ListSecurityDataClassificationRulesRequestOrderByEnum {
	return ListSecurityDataClassificationRulesRequestOrderByEnum{
		CREATED_AT: ListSecurityDataClassificationRulesRequestOrderBy{
			value: "createdAt",
		},
		CREATED_BY: ListSecurityDataClassificationRulesRequestOrderBy{
			value: "createdBy",
		},
		UPDATED_AT: ListSecurityDataClassificationRulesRequestOrderBy{
			value: "updatedAt",
		},
		UPDATED_BY: ListSecurityDataClassificationRulesRequestOrderBy{
			value: "updatedBy",
		},
		NAME: ListSecurityDataClassificationRulesRequestOrderBy{
			value: "name",
		},
		DESCRIPTION: ListSecurityDataClassificationRulesRequestOrderBy{
			value: "description",
		},
	}
}

func (c ListSecurityDataClassificationRulesRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityDataClassificationRulesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityDataClassificationRulesRequestOrderBy) UnmarshalJSON(b []byte) error {
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
