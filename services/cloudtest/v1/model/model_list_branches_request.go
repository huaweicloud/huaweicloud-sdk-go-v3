package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBranchesRequest Request Object
type ListBranchesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 排序字段
	SortField *ListBranchesRequestSortField `json:"sort_field,omitempty"`

	// 排序方式
	SortType *ListBranchesRequestSortType `json:"sort_type,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询， offset大于等于0，小于等于20000
	Offset int32 `json:"offset"`

	// 每页显示的条目数量，最大支持200条
	Limit int32 `json:"limit"`
}

func (o ListBranchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListBranchesRequest", string(data)}, " ")
}

type ListBranchesRequestSortField struct {
	value string
}

type ListBranchesRequestSortFieldEnum struct {
	NAME          ListBranchesRequestSortField
	CREATION_DATE ListBranchesRequestSortField
}

func GetListBranchesRequestSortFieldEnum() ListBranchesRequestSortFieldEnum {
	return ListBranchesRequestSortFieldEnum{
		NAME: ListBranchesRequestSortField{
			value: "name",
		},
		CREATION_DATE: ListBranchesRequestSortField{
			value: "creationDate",
		},
	}
}

func (c ListBranchesRequestSortField) Value() string {
	return c.value
}

func (c ListBranchesRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBranchesRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListBranchesRequestSortType struct {
	value string
}

type ListBranchesRequestSortTypeEnum struct {
	ASC  ListBranchesRequestSortType
	DESC ListBranchesRequestSortType
}

func GetListBranchesRequestSortTypeEnum() ListBranchesRequestSortTypeEnum {
	return ListBranchesRequestSortTypeEnum{
		ASC: ListBranchesRequestSortType{
			value: "ASC",
		},
		DESC: ListBranchesRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListBranchesRequestSortType) Value() string {
	return c.value
}

func (c ListBranchesRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBranchesRequestSortType) UnmarshalJSON(b []byte) error {
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
