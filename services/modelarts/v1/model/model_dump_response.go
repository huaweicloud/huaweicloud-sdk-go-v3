package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DumpResponse **参数解释：** 用户转储配置。
type DumpResponse struct {

	// **参数解释：** 用户转储来源类别。 **取值范围：** - [OBS：对象存储服务。](tag:hws,hws_hk) - OBSFS：OBS的文件系统接口。
	Source DumpResponseSource `json:"source"`

	// **参数解释：** 用户转储来源地址，格式遵循不同存储系统。 **取值范围：** 不涉及。
	Address *string `json:"address,omitempty"`

	// **参数解释：** 挂载到容器内的路径，要求以/开头，后面可包含中划线，反斜杠，下划线，点号，字母，数字。 **取值范围：** 不涉及。
	MountPath string `json:"mount_path"`

	// **参数解释：** 转储类型。 **取值范围：** - DUMP：用户自定义转储。 - DUMP_SYS：系统转储。
	Type *DumpResponseType `json:"type,omitempty"`
}

func (o DumpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DumpResponse struct{}"
	}

	return strings.Join([]string{"DumpResponse", string(data)}, " ")
}

type DumpResponseSource struct {
	value string
}

type DumpResponseSourceEnum struct {
	OBS   DumpResponseSource
	OBSFS DumpResponseSource
}

func GetDumpResponseSourceEnum() DumpResponseSourceEnum {
	return DumpResponseSourceEnum{
		OBS: DumpResponseSource{
			value: "OBS",
		},
		OBSFS: DumpResponseSource{
			value: "OBSFS",
		},
	}
}

func (c DumpResponseSource) Value() string {
	return c.value
}

func (c DumpResponseSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DumpResponseSource) UnmarshalJSON(b []byte) error {
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

type DumpResponseType struct {
	value string
}

type DumpResponseTypeEnum struct {
	DUMP     DumpResponseType
	DUMP_SYS DumpResponseType
}

func GetDumpResponseTypeEnum() DumpResponseTypeEnum {
	return DumpResponseTypeEnum{
		DUMP: DumpResponseType{
			value: "DUMP",
		},
		DUMP_SYS: DumpResponseType{
			value: "DUMP_SYS",
		},
	}
}

func (c DumpResponseType) Value() string {
	return c.value
}

func (c DumpResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DumpResponseType) UnmarshalJSON(b []byte) error {
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
