package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecuritySensitiveDataDetailsRequest Request Object
type ListSecuritySensitiveDataDetailsRequest struct {

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 用于游标分页，表示查询ID大于该值的记录（不包含该ID），仅支持向前翻页，且不可与offset参数同时使用。
	Marker *string `json:"marker,omitempty"`

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 敏感数据发现开始时间。
	FindStartTime *int64 `json:"find_start_time,omitempty"`

	// 敏感数据发现结束时间。
	FindEndTime *int64 `json:"find_end_time,omitempty"`

	// 排序字段，FIND_TIME（仅使用limit、offset分页时有效）。
	OrderBy *ListSecuritySensitiveDataDetailsRequestOrderBy `json:"order_by,omitempty"`

	// 是否升序（仅指定排序参数，且使用limit、offset分页时有效）。
	OrderByAsc *bool `json:"order_by_asc,omitempty"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ListSecuritySensitiveDataDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecuritySensitiveDataDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSecuritySensitiveDataDetailsRequest", string(data)}, " ")
}

type ListSecuritySensitiveDataDetailsRequestOrderBy struct {
	value string
}

type ListSecuritySensitiveDataDetailsRequestOrderByEnum struct {
	FIND_TIME ListSecuritySensitiveDataDetailsRequestOrderBy
}

func GetListSecuritySensitiveDataDetailsRequestOrderByEnum() ListSecuritySensitiveDataDetailsRequestOrderByEnum {
	return ListSecuritySensitiveDataDetailsRequestOrderByEnum{
		FIND_TIME: ListSecuritySensitiveDataDetailsRequestOrderBy{
			value: "FIND_TIME",
		},
	}
}

func (c ListSecuritySensitiveDataDetailsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecuritySensitiveDataDetailsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecuritySensitiveDataDetailsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
