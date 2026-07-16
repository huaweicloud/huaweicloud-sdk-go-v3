package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DevServerBatchRequest Lite Server实例批量操作请求体。
type DevServerBatchRequest struct {

	// **参数解释**：批量操作类型。 **约束限制**：不涉及。 **取值范围**： - START：批量启动Lite Server实例 - STOP：批量停止Lite Server实例 - REBOOT：批量重启Lite Server实例 - CHANGEOS：批量切换Lite Server服务器操作系统镜像 - REINSTALLOS：批量重装Lite Server服务器操作系统镜像 - DELETE：批量删除Lite Server实例 **默认取值**：不涉及。
	Type DevServerBatchRequestType `json:"type"`

	// **参数解释**：批量操作Lite Server ID列表。
	Servers []BatchActionDevServerIds `json:"servers"`

	ExtendParam *ServerOsRequest `json:"extend_param,omitempty"`
}

func (o DevServerBatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevServerBatchRequest struct{}"
	}

	return strings.Join([]string{"DevServerBatchRequest", string(data)}, " ")
}

type DevServerBatchRequestType struct {
	value string
}

type DevServerBatchRequestTypeEnum struct {
	START       DevServerBatchRequestType
	STOP        DevServerBatchRequestType
	REBOOT      DevServerBatchRequestType
	CHANGEOS    DevServerBatchRequestType
	REINSTALLOS DevServerBatchRequestType
	DELETE      DevServerBatchRequestType
}

func GetDevServerBatchRequestTypeEnum() DevServerBatchRequestTypeEnum {
	return DevServerBatchRequestTypeEnum{
		START: DevServerBatchRequestType{
			value: "START",
		},
		STOP: DevServerBatchRequestType{
			value: "STOP",
		},
		REBOOT: DevServerBatchRequestType{
			value: "REBOOT",
		},
		CHANGEOS: DevServerBatchRequestType{
			value: "CHANGEOS",
		},
		REINSTALLOS: DevServerBatchRequestType{
			value: "REINSTALLOS",
		},
		DELETE: DevServerBatchRequestType{
			value: "DELETE",
		},
	}
}

func (c DevServerBatchRequestType) Value() string {
	return c.value
}

func (c DevServerBatchRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevServerBatchRequestType) UnmarshalJSON(b []byte) error {
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
