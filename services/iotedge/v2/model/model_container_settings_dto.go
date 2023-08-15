package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ContainerSettingsDto struct {
	Configs *ContainerConfigsDto `json:"configs,omitempty"`

	// 镜像存储地址
	ImageUrl string `json:"image_url"`

	// 环境变量
	Envs *interface{} `json:"envs,omitempty"`

	// 卷配置
	Volumes *[]VolumeDto `json:"volumes,omitempty"`

	// NPU类型, D310:昇腾310推理卡，D910:昇腾910训练卡;D310P：昇腾710或者310P加速卡
	NpuType *ContainerSettingsDtoNpuType `json:"npu_type,omitempty"`

	// NPU算力切分模板,昇腾D310Pro，支持：vir01、vir02、vir02_1c、vir04、vir04_4c_dvpp、vir04_3c、vir04_3c_ndvpp 昇腾D910芯片支持:vir01|vir02|vir04|vir08 可在对应芯片的机器上通过npu-smi info -t template-info命令查询其详细信息
	VnpuTemplate *ContainerSettingsDtoVnpuTemplate `json:"vnpu_template,omitempty"`

	Resources *ResourceDto `json:"resources,omitempty"`

	// 外挂设备配置
	ExtDevices *[]ExtDevice `json:"ext_devices,omitempty"`
}

func (o ContainerSettingsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerSettingsDto struct{}"
	}

	return strings.Join([]string{"ContainerSettingsDto", string(data)}, " ")
}

type ContainerSettingsDtoNpuType struct {
	value string
}

type ContainerSettingsDtoNpuTypeEnum struct {
	D310   ContainerSettingsDtoNpuType
	D910   ContainerSettingsDtoNpuType
	D310_P ContainerSettingsDtoNpuType
}

func GetContainerSettingsDtoNpuTypeEnum() ContainerSettingsDtoNpuTypeEnum {
	return ContainerSettingsDtoNpuTypeEnum{
		D310: ContainerSettingsDtoNpuType{
			value: "D310",
		},
		D910: ContainerSettingsDtoNpuType{
			value: "D910",
		},
		D310_P: ContainerSettingsDtoNpuType{
			value: "D310P",
		},
	}
}

func (c ContainerSettingsDtoNpuType) Value() string {
	return c.value
}

func (c ContainerSettingsDtoNpuType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContainerSettingsDtoNpuType) UnmarshalJSON(b []byte) error {
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

type ContainerSettingsDtoVnpuTemplate struct {
	value string
}

type ContainerSettingsDtoVnpuTemplateEnum struct {
	VIR01          ContainerSettingsDtoVnpuTemplate
	VIR02          ContainerSettingsDtoVnpuTemplate
	VIR04          ContainerSettingsDtoVnpuTemplate
	VIR08          ContainerSettingsDtoVnpuTemplate
	VIR02_1C       ContainerSettingsDtoVnpuTemplate
	VIR04_3C       ContainerSettingsDtoVnpuTemplate
	VIR04_3C_NDVPP ContainerSettingsDtoVnpuTemplate
	VIR04_4C_DVPP  ContainerSettingsDtoVnpuTemplate
}

func GetContainerSettingsDtoVnpuTemplateEnum() ContainerSettingsDtoVnpuTemplateEnum {
	return ContainerSettingsDtoVnpuTemplateEnum{
		VIR01: ContainerSettingsDtoVnpuTemplate{
			value: "vir01",
		},
		VIR02: ContainerSettingsDtoVnpuTemplate{
			value: "vir02",
		},
		VIR04: ContainerSettingsDtoVnpuTemplate{
			value: "vir04",
		},
		VIR08: ContainerSettingsDtoVnpuTemplate{
			value: "vir08",
		},
		VIR02_1C: ContainerSettingsDtoVnpuTemplate{
			value: "vir02_1c",
		},
		VIR04_3C: ContainerSettingsDtoVnpuTemplate{
			value: "vir04_3c",
		},
		VIR04_3C_NDVPP: ContainerSettingsDtoVnpuTemplate{
			value: "vir04_3c_ndvpp",
		},
		VIR04_4C_DVPP: ContainerSettingsDtoVnpuTemplate{
			value: "vir04_4c_dvpp",
		},
	}
}

func (c ContainerSettingsDtoVnpuTemplate) Value() string {
	return c.value
}

func (c ContainerSettingsDtoVnpuTemplate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContainerSettingsDtoVnpuTemplate) UnmarshalJSON(b []byte) error {
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
