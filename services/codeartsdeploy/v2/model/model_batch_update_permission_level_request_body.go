package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdatePermissionLevelRequestBody 批量配置应用下鉴权级别请求体
type BatchUpdatePermissionLevelRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 应用鉴权级别，instance：实例级；project：项目级
	PermissionLevel BatchUpdatePermissionLevelRequestBodyPermissionLevel `json:"permission_level"`

	// 应用id列表
	ApplicationIds []string `json:"application_ids"`
}

func (o BatchUpdatePermissionLevelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePermissionLevelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePermissionLevelRequestBody", string(data)}, " ")
}

type BatchUpdatePermissionLevelRequestBodyPermissionLevel struct {
	value string
}

type BatchUpdatePermissionLevelRequestBodyPermissionLevelEnum struct {
	PROJECT  BatchUpdatePermissionLevelRequestBodyPermissionLevel
	INSTANCE BatchUpdatePermissionLevelRequestBodyPermissionLevel
}

func GetBatchUpdatePermissionLevelRequestBodyPermissionLevelEnum() BatchUpdatePermissionLevelRequestBodyPermissionLevelEnum {
	return BatchUpdatePermissionLevelRequestBodyPermissionLevelEnum{
		PROJECT: BatchUpdatePermissionLevelRequestBodyPermissionLevel{
			value: "project",
		},
		INSTANCE: BatchUpdatePermissionLevelRequestBodyPermissionLevel{
			value: "instance",
		},
	}
}

func (c BatchUpdatePermissionLevelRequestBodyPermissionLevel) Value() string {
	return c.value
}

func (c BatchUpdatePermissionLevelRequestBodyPermissionLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdatePermissionLevelRequestBodyPermissionLevel) UnmarshalJSON(b []byte) error {
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
