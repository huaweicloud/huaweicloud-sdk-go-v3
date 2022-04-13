package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListEnterpriseProjectRequest struct {
	// 企业项目ID，0表示默认企业项目

	Id *string `json:"id,omitempty"`
	// 查询记录数默认为1000，limit最多为1000, 最小值为1

	Limit *int32 `json:"limit,omitempty"`
	// 企业项目名称，支持模糊搜索

	Name *string `json:"name,omitempty"`
	// 索引位置，从offset指定的下一条数据开始查询，必须为数字，不能为负数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 降序或升序,默认为“desc” 。desc表示降序 。asc 表示升序

	SortDir *ListEnterpriseProjectRequestSortDir `json:"sort_dir,omitempty"`
	// 返回结果按该关键字排序（支持updated_at等关键字，默认为“created_at”）

	SortKey *ListEnterpriseProjectRequestSortKey `json:"sort_key,omitempty"`
	// 企业项目状态。 1--启用，2--停用

	Status *int32 `json:"status,omitempty"`
}

func (o ListEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectRequest", string(data)}, " ")
}

type ListEnterpriseProjectRequestSortDir struct {
	value string
}

type ListEnterpriseProjectRequestSortDirEnum struct {
	DESC ListEnterpriseProjectRequestSortDir
	ASC  ListEnterpriseProjectRequestSortDir
}

func GetListEnterpriseProjectRequestSortDirEnum() ListEnterpriseProjectRequestSortDirEnum {
	return ListEnterpriseProjectRequestSortDirEnum{
		DESC: ListEnterpriseProjectRequestSortDir{
			value: "desc",
		},
		ASC: ListEnterpriseProjectRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListEnterpriseProjectRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnterpriseProjectRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListEnterpriseProjectRequestSortKey struct {
	value string
}

type ListEnterpriseProjectRequestSortKeyEnum struct {
	CREATED_AT ListEnterpriseProjectRequestSortKey
	UPDATED_AT ListEnterpriseProjectRequestSortKey
}

func GetListEnterpriseProjectRequestSortKeyEnum() ListEnterpriseProjectRequestSortKeyEnum {
	return ListEnterpriseProjectRequestSortKeyEnum{
		CREATED_AT: ListEnterpriseProjectRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListEnterpriseProjectRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListEnterpriseProjectRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnterpriseProjectRequestSortKey) UnmarshalJSON(b []byte) error {
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
