package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceGroupsRequest Request Object
type ListResourceGroupsRequest struct {

	// 归属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源分组名称，支持模糊查询
	GroupName *string `json:"group_name,omitempty"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId *string `json:"group_id,omitempty"`

	// 分页查询时查询的起始位置，表示从第几条数据开始，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询时每页的条目数，取值[1,100]，默认值为100
	Limit *int32 `json:"limit,omitempty"`

	// 资源分组添加资源方式，取值只能为EPS（同步企业项目）,TAG（标签动态匹配）,Manual（手动添加），不传代表查询所有资源分组类型
	Type *ListResourceGroupsRequestType `json:"type,omitempty"`
}

func (o ListResourceGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupsRequest", string(data)}, " ")
}

type ListResourceGroupsRequestType struct {
	value string
}

type ListResourceGroupsRequestTypeEnum struct {
	EPS    ListResourceGroupsRequestType
	TAG    ListResourceGroupsRequestType
	MANUAL ListResourceGroupsRequestType
}

func GetListResourceGroupsRequestTypeEnum() ListResourceGroupsRequestTypeEnum {
	return ListResourceGroupsRequestTypeEnum{
		EPS: ListResourceGroupsRequestType{
			value: "EPS",
		},
		TAG: ListResourceGroupsRequestType{
			value: "TAG",
		},
		MANUAL: ListResourceGroupsRequestType{
			value: "Manual",
		},
	}
}

func (c ListResourceGroupsRequestType) Value() string {
	return c.value
}

func (c ListResourceGroupsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupsRequestType) UnmarshalJSON(b []byte) error {
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
