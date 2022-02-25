package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchAddSharedTagsRequestBody struct {
	// 操作标识，取值范围为：\"create\"。 为指定共享批量添加标签时使用\"create\"。

	Action BatchAddSharedTagsRequestBodyAction `json:"action"`
	// 标签列表。 用户权限时该字段必选，op_service权限时和sys_tags二选一。

	Tags *[]ResourceTag `json:"tags,omitempty"`
	// 系统标签列表。 op_service权限可以访问，和tags二选一。目前TMS调用时只包含一个resource_tag结构体，key固定为：_sys_enterprise_project_id。

	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`
}

func (o BatchAddSharedTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsRequestBody", string(data)}, " ")
}

type BatchAddSharedTagsRequestBodyAction struct {
	value string
}

type BatchAddSharedTagsRequestBodyActionEnum struct {
	CREATE BatchAddSharedTagsRequestBodyAction
}

func GetBatchAddSharedTagsRequestBodyActionEnum() BatchAddSharedTagsRequestBodyActionEnum {
	return BatchAddSharedTagsRequestBodyActionEnum{
		CREATE: BatchAddSharedTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchAddSharedTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddSharedTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
