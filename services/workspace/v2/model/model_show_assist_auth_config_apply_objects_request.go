package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAssistAuthConfigApplyObjectsRequest Request Object
type ShowAssistAuthConfigApplyObjectsRequest struct {

	// 绑定对象类型枚举。 - USER：用户 - USER_GROUP：用户组 - ALL: 全部用户
	ObjectType *ShowAssistAuthConfigApplyObjectsRequestObjectType `json:"object_type,omitempty"`

	// 对象名称。
	ObjectName *string `json:"object_name,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回查询应用对象限制。取值范围1-1000，默认值是100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowAssistAuthConfigApplyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssistAuthConfigApplyObjectsRequest struct{}"
	}

	return strings.Join([]string{"ShowAssistAuthConfigApplyObjectsRequest", string(data)}, " ")
}

type ShowAssistAuthConfigApplyObjectsRequestObjectType struct {
	value string
}

type ShowAssistAuthConfigApplyObjectsRequestObjectTypeEnum struct {
	USER       ShowAssistAuthConfigApplyObjectsRequestObjectType
	USER_GROUP ShowAssistAuthConfigApplyObjectsRequestObjectType
	ALL        ShowAssistAuthConfigApplyObjectsRequestObjectType
}

func GetShowAssistAuthConfigApplyObjectsRequestObjectTypeEnum() ShowAssistAuthConfigApplyObjectsRequestObjectTypeEnum {
	return ShowAssistAuthConfigApplyObjectsRequestObjectTypeEnum{
		USER: ShowAssistAuthConfigApplyObjectsRequestObjectType{
			value: "USER",
		},
		USER_GROUP: ShowAssistAuthConfigApplyObjectsRequestObjectType{
			value: "USER_GROUP",
		},
		ALL: ShowAssistAuthConfigApplyObjectsRequestObjectType{
			value: "ALL",
		},
	}
}

func (c ShowAssistAuthConfigApplyObjectsRequestObjectType) Value() string {
	return c.value
}

func (c ShowAssistAuthConfigApplyObjectsRequestObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAssistAuthConfigApplyObjectsRequestObjectType) UnmarshalJSON(b []byte) error {
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
