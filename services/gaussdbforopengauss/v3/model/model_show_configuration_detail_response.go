package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowConfigurationDetailResponse struct {

	// 参数模板ID。
	Id *string `json:"id,omitempty"`

	// 参数模板名称。
	Name *string `json:"name,omitempty"`

	// 参数模板描述。
	Description *string `json:"description,omitempty"`

	// 引擎版本。 [数据库版本。支持2.3版本，取值为“2.3”]。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例部署形态。independent：独立；ha：主备。
	InstanceMode *ShowConfigurationDetailResponseInstanceMode `json:"instance_mode,omitempty"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 参数详情。
	ConfigurationParameters *[]ParaGroupParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                         `json:"-"`
}

func (o ShowConfigurationDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailResponse", string(data)}, " ")
}

type ShowConfigurationDetailResponseInstanceMode struct {
	value string
}

type ShowConfigurationDetailResponseInstanceModeEnum struct {
	INDEPENDENT ShowConfigurationDetailResponseInstanceMode
	HA          ShowConfigurationDetailResponseInstanceMode
}

func GetShowConfigurationDetailResponseInstanceModeEnum() ShowConfigurationDetailResponseInstanceModeEnum {
	return ShowConfigurationDetailResponseInstanceModeEnum{
		INDEPENDENT: ShowConfigurationDetailResponseInstanceMode{
			value: "independent",
		},
		HA: ShowConfigurationDetailResponseInstanceMode{
			value: "ha",
		},
	}
}

func (c ShowConfigurationDetailResponseInstanceMode) Value() string {
	return c.value
}

func (c ShowConfigurationDetailResponseInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigurationDetailResponseInstanceMode) UnmarshalJSON(b []byte) error {
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
