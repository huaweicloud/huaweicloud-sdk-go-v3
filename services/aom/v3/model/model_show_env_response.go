package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEnvResponse Response Object
type ShowEnvResponse struct {

	// aomId
	AomId *string `json:"aom_id,omitempty"`

	// 组件Id
	ComponentId *string `json:"component_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 环境Id
	EnvId *string `json:"env_id,omitempty"`

	// 环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 环境标签
	EnvTags *[]TagNameAndIdVo `json:"env_tags,omitempty"`

	// 环境类型
	EnvType *string `json:"env_type,omitempty"`

	// 企业项目Id
	EpsId *string `json:"eps_id,omitempty"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 修改者
	Modifier *string `json:"modifier,omitempty"`

	// os类型
	OsType *string `json:"os_type,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 注册方式
	RegisterType   *ShowEnvResponseRegisterType `json:"register_type,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvResponse struct{}"
	}

	return strings.Join([]string{"ShowEnvResponse", string(data)}, " ")
}

type ShowEnvResponseRegisterType struct {
	value string
}

type ShowEnvResponseRegisterTypeEnum struct {
	API               ShowEnvResponseRegisterType
	CONSOLE           ShowEnvResponseRegisterType
	SERVICE_DISCOVERY ShowEnvResponseRegisterType
}

func GetShowEnvResponseRegisterTypeEnum() ShowEnvResponseRegisterTypeEnum {
	return ShowEnvResponseRegisterTypeEnum{
		API: ShowEnvResponseRegisterType{
			value: "API",
		},
		CONSOLE: ShowEnvResponseRegisterType{
			value: "CONSOLE",
		},
		SERVICE_DISCOVERY: ShowEnvResponseRegisterType{
			value: "SERVICE_DISCOVERY",
		},
	}
}

func (c ShowEnvResponseRegisterType) Value() string {
	return c.value
}

func (c ShowEnvResponseRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEnvResponseRegisterType) UnmarshalJSON(b []byte) error {
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
