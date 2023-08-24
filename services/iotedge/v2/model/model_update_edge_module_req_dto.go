package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEdgeModuleReqDto 更新边缘模块请求结构体
type UpdateEdgeModuleReqDto struct {

	// 边缘应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 边缘模块名称
	ModuleName *string `json:"module_name,omitempty"`

	ContainerSettings *ContainerSettingsReqDto `json:"container_settings,omitempty"`

	// 模块期望状态: RUNNING(升级后期望模块运行)，STOPPED(升级后期望模块停止)，空值默认继承升级前模块期望状态
	DesiredState *UpdateEdgeModuleReqDtoDesiredState `json:"desired_state,omitempty"`
}

func (o UpdateEdgeModuleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeModuleReqDto", string(data)}, " ")
}

type UpdateEdgeModuleReqDtoDesiredState struct {
	value string
}

type UpdateEdgeModuleReqDtoDesiredStateEnum struct {
	RUNNING UpdateEdgeModuleReqDtoDesiredState
	STOPPED UpdateEdgeModuleReqDtoDesiredState
}

func GetUpdateEdgeModuleReqDtoDesiredStateEnum() UpdateEdgeModuleReqDtoDesiredStateEnum {
	return UpdateEdgeModuleReqDtoDesiredStateEnum{
		RUNNING: UpdateEdgeModuleReqDtoDesiredState{
			value: "RUNNING",
		},
		STOPPED: UpdateEdgeModuleReqDtoDesiredState{
			value: "STOPPED",
		},
	}
}

func (c UpdateEdgeModuleReqDtoDesiredState) Value() string {
	return c.value
}

func (c UpdateEdgeModuleReqDtoDesiredState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeModuleReqDtoDesiredState) UnmarshalJSON(b []byte) error {
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
