package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowParameterGroupDetailResponse Response Object
type ShowParameterGroupDetailResponse struct {

	// 参数模板ID。
	Id *string `json:"id,omitempty"`

	// 参数模板名称。
	Name *string `json:"name,omitempty"`

	// 参数模板描述。
	Description *string `json:"description,omitempty"`

	// 引擎版本。 [数据库版本。支持V2.0-2.3版本，取值为“V2.0-2.3”]。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例部署形态。independent：独立；ha：主备。
	InstanceMode *ShowParameterGroupDetailResponseInstanceMode `json:"instance_mode,omitempty"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 参数详情。
	ConfigurationParameters *[]ParaGroupParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                         `json:"-"`
}

func (o ShowParameterGroupDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowParameterGroupDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowParameterGroupDetailResponse", string(data)}, " ")
}

type ShowParameterGroupDetailResponseInstanceMode struct {
	value string
}

type ShowParameterGroupDetailResponseInstanceModeEnum struct {
	INDEPENDENT ShowParameterGroupDetailResponseInstanceMode
	HA          ShowParameterGroupDetailResponseInstanceMode
}

func GetShowParameterGroupDetailResponseInstanceModeEnum() ShowParameterGroupDetailResponseInstanceModeEnum {
	return ShowParameterGroupDetailResponseInstanceModeEnum{
		INDEPENDENT: ShowParameterGroupDetailResponseInstanceMode{
			value: "independent",
		},
		HA: ShowParameterGroupDetailResponseInstanceMode{
			value: "ha",
		},
	}
}

func (c ShowParameterGroupDetailResponseInstanceMode) Value() string {
	return c.value
}

func (c ShowParameterGroupDetailResponseInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowParameterGroupDetailResponseInstanceMode) UnmarshalJSON(b []byte) error {
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
