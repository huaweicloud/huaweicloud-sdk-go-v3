package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAppInfoRequestBody 更新应用请求体
type UpdateAppInfoRequestBody struct {

	// 应用id
	Id string `json:"id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 应用名称
	Name string `json:"name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否为草稿
	IsDraft bool `json:"is_draft"`

	// 创建类型，template：部署模板创建
	CreateType UpdateAppInfoRequestBodyCreateType `json:"create_type"`

	// 自定义slave资源池id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	Trigger *TaskTriggerVo `json:"trigger,omitempty"`

	// 部署任务列表信息
	ArrangeInfos *[]UpdateTaskV2RequestBody `json:"arrange_infos,omitempty"`
}

func (o UpdateAppInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppInfoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAppInfoRequestBody", string(data)}, " ")
}

type UpdateAppInfoRequestBodyCreateType struct {
	value string
}

type UpdateAppInfoRequestBodyCreateTypeEnum struct {
	TEMPLATE UpdateAppInfoRequestBodyCreateType
}

func GetUpdateAppInfoRequestBodyCreateTypeEnum() UpdateAppInfoRequestBodyCreateTypeEnum {
	return UpdateAppInfoRequestBodyCreateTypeEnum{
		TEMPLATE: UpdateAppInfoRequestBodyCreateType{
			value: "template",
		},
	}
}

func (c UpdateAppInfoRequestBodyCreateType) Value() string {
	return c.value
}

func (c UpdateAppInfoRequestBodyCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAppInfoRequestBodyCreateType) UnmarshalJSON(b []byte) error {
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
