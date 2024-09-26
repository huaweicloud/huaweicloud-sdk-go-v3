package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListManageableGroupsRequest Request Object
type ListManageableGroupsRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 创建资源类型，group 代码组，repository仓库
	Scope *ListManageableGroupsRequestScope `json:"scope,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 返回数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListManageableGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManageableGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListManageableGroupsRequest", string(data)}, " ")
}

type ListManageableGroupsRequestScope struct {
	value string
}

type ListManageableGroupsRequestScopeEnum struct {
	GROUP      ListManageableGroupsRequestScope
	REPOSITORY ListManageableGroupsRequestScope
}

func GetListManageableGroupsRequestScopeEnum() ListManageableGroupsRequestScopeEnum {
	return ListManageableGroupsRequestScopeEnum{
		GROUP: ListManageableGroupsRequestScope{
			value: "group",
		},
		REPOSITORY: ListManageableGroupsRequestScope{
			value: "repository",
		},
	}
}

func (c ListManageableGroupsRequestScope) Value() string {
	return c.value
}

func (c ListManageableGroupsRequestScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListManageableGroupsRequestScope) UnmarshalJSON(b []byte) error {
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
