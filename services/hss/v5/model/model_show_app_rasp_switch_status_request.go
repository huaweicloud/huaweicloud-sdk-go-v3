package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppRaspSwitchStatusRequest Request Object
type ShowAppRaspSwitchStatusRequest struct {

	// 企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 应用类型，包含如下1种。   - java ：java类型应用防护。
	AppType *ShowAppRaspSwitchStatusRequestAppType `json:"app_type,omitempty"`

	// host id
	HostId string `json:"host_id"`
}

func (o ShowAppRaspSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppRaspSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAppRaspSwitchStatusRequest", string(data)}, " ")
}

type ShowAppRaspSwitchStatusRequestAppType struct {
	value string
}

type ShowAppRaspSwitchStatusRequestAppTypeEnum struct {
	JAVA ShowAppRaspSwitchStatusRequestAppType
}

func GetShowAppRaspSwitchStatusRequestAppTypeEnum() ShowAppRaspSwitchStatusRequestAppTypeEnum {
	return ShowAppRaspSwitchStatusRequestAppTypeEnum{
		JAVA: ShowAppRaspSwitchStatusRequestAppType{
			value: "java",
		},
	}
}

func (c ShowAppRaspSwitchStatusRequestAppType) Value() string {
	return c.value
}

func (c ShowAppRaspSwitchStatusRequestAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppRaspSwitchStatusRequestAppType) UnmarshalJSON(b []byte) error {
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
