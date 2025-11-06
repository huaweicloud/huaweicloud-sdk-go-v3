package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupPermissionResourcesRequest Request Object
type ListGroupPermissionResourcesRequest struct {

	// **参数解释：** 创建资源类型。 **约束限制：** - group 代码组。 - project  项目。 - all 代码组和项目
	Scope *ListGroupPermissionResourcesRequestScope `json:"scope,omitempty"`
}

func (o ListGroupPermissionResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupPermissionResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListGroupPermissionResourcesRequest", string(data)}, " ")
}

type ListGroupPermissionResourcesRequestScope struct {
	value string
}

type ListGroupPermissionResourcesRequestScopeEnum struct {
	GROUP   ListGroupPermissionResourcesRequestScope
	PROJECT ListGroupPermissionResourcesRequestScope
	ALL     ListGroupPermissionResourcesRequestScope
}

func GetListGroupPermissionResourcesRequestScopeEnum() ListGroupPermissionResourcesRequestScopeEnum {
	return ListGroupPermissionResourcesRequestScopeEnum{
		GROUP: ListGroupPermissionResourcesRequestScope{
			value: "group",
		},
		PROJECT: ListGroupPermissionResourcesRequestScope{
			value: "project",
		},
		ALL: ListGroupPermissionResourcesRequestScope{
			value: "all",
		},
	}
}

func (c ListGroupPermissionResourcesRequestScope) Value() string {
	return c.value
}

func (c ListGroupPermissionResourcesRequestScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupPermissionResourcesRequestScope) UnmarshalJSON(b []byte) error {
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
