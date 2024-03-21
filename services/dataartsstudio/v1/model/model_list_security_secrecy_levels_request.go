package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecuritySecrecyLevelsRequest Request Object
type ListSecuritySecrecyLevelsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段, createdAt, createdBy, updatedAt, updatedBy, name, description
	OrderBy *ListSecuritySecrecyLevelsRequestOrderBy `json:"order_by,omitempty"`

	// 排序规则
	Desc *bool `json:"desc,omitempty"`
}

func (o ListSecuritySecrecyLevelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecuritySecrecyLevelsRequest struct{}"
	}

	return strings.Join([]string{"ListSecuritySecrecyLevelsRequest", string(data)}, " ")
}

type ListSecuritySecrecyLevelsRequestOrderBy struct {
	value string
}

type ListSecuritySecrecyLevelsRequestOrderByEnum struct {
	CREATED_AT  ListSecuritySecrecyLevelsRequestOrderBy
	CREATED_BY  ListSecuritySecrecyLevelsRequestOrderBy
	UPDATED_AT  ListSecuritySecrecyLevelsRequestOrderBy
	UPDATED_BY  ListSecuritySecrecyLevelsRequestOrderBy
	NAME        ListSecuritySecrecyLevelsRequestOrderBy
	DESCRIPTION ListSecuritySecrecyLevelsRequestOrderBy
}

func GetListSecuritySecrecyLevelsRequestOrderByEnum() ListSecuritySecrecyLevelsRequestOrderByEnum {
	return ListSecuritySecrecyLevelsRequestOrderByEnum{
		CREATED_AT: ListSecuritySecrecyLevelsRequestOrderBy{
			value: "createdAt",
		},
		CREATED_BY: ListSecuritySecrecyLevelsRequestOrderBy{
			value: "createdBy",
		},
		UPDATED_AT: ListSecuritySecrecyLevelsRequestOrderBy{
			value: "updatedAt",
		},
		UPDATED_BY: ListSecuritySecrecyLevelsRequestOrderBy{
			value: "updatedBy",
		},
		NAME: ListSecuritySecrecyLevelsRequestOrderBy{
			value: "name",
		},
		DESCRIPTION: ListSecuritySecrecyLevelsRequestOrderBy{
			value: "description",
		},
	}
}

func (c ListSecuritySecrecyLevelsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecuritySecrecyLevelsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecuritySecrecyLevelsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
