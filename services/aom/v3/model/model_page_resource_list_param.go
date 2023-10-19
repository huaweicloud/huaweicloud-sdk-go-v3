package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PageResourceListParam struct {

	// 页面的分页标志位,为分页的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 查询返回记录的数量限制。limit可以为空，如果值小于1或者大于100，则会使用默认值100
	Limit *int32 `json:"limit,omitempty"`

	// 关键字模糊搜索。Key取值：NAME、RESOURCE_ID
	Keywords map[string]string `json:"keywords,omitempty"`

	// 是否需要返回拓扑树,默认是false。需要：true---性能差，不需要false--性能好
	CiRelationships *bool `json:"ci_relationships,omitempty"`

	// 节点类型，取值：application、sub_application、component、environment
	CiType PageResourceListParamCiType `json:"ci_type"`

	// 环境的region信息，若没有值，代表全部
	CiRegion *string `json:"ci_region,omitempty"`

	// 节点id列表;如果ci_ids和ci_id同时有，则优先ci_ids，但是不能同时为空
	CiIds *[]string `json:"ci_ids,omitempty"`

	// 节点id列表;如果ci_ids和ci_id同时有，则优先ci_ids，但是不能同时为空。但是不支持应用批量查询
	CiId *string `json:"ci_id,omitempty"`
}

func (o PageResourceListParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageResourceListParam struct{}"
	}

	return strings.Join([]string{"PageResourceListParam", string(data)}, " ")
}

type PageResourceListParamCiType struct {
	value string
}

type PageResourceListParamCiTypeEnum struct {
	APPLICATION     PageResourceListParamCiType
	SUB_APPLICATION PageResourceListParamCiType
	COMPONENT       PageResourceListParamCiType
	ENVIRONMENT     PageResourceListParamCiType
}

func GetPageResourceListParamCiTypeEnum() PageResourceListParamCiTypeEnum {
	return PageResourceListParamCiTypeEnum{
		APPLICATION: PageResourceListParamCiType{
			value: "APPLICATION",
		},
		SUB_APPLICATION: PageResourceListParamCiType{
			value: "SUB_APPLICATION",
		},
		COMPONENT: PageResourceListParamCiType{
			value: "COMPONENT",
		},
		ENVIRONMENT: PageResourceListParamCiType{
			value: "ENVIRONMENT",
		},
	}
}

func (c PageResourceListParamCiType) Value() string {
	return c.value
}

func (c PageResourceListParamCiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PageResourceListParamCiType) UnmarshalJSON(b []byte) error {
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
