package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDataServiceInstanceObsLogRequest Request Object
type UpdateDataServiceInstanceObsLogRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *UpdateDataServiceInstanceObsLogRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType *string `json:"Content-Type,omitempty"`

	// 集群ID编号。
	InstanceId string `json:"instance_id"`

	Body *ObsLogDump `json:"body,omitempty"`
}

func (o UpdateDataServiceInstanceObsLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataServiceInstanceObsLogRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataServiceInstanceObsLogRequest", string(data)}, " ")
}

type UpdateDataServiceInstanceObsLogRequestDlmType struct {
	value string
}

type UpdateDataServiceInstanceObsLogRequestDlmTypeEnum struct {
	SHARED    UpdateDataServiceInstanceObsLogRequestDlmType
	EXCLUSIVE UpdateDataServiceInstanceObsLogRequestDlmType
}

func GetUpdateDataServiceInstanceObsLogRequestDlmTypeEnum() UpdateDataServiceInstanceObsLogRequestDlmTypeEnum {
	return UpdateDataServiceInstanceObsLogRequestDlmTypeEnum{
		SHARED: UpdateDataServiceInstanceObsLogRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: UpdateDataServiceInstanceObsLogRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c UpdateDataServiceInstanceObsLogRequestDlmType) Value() string {
	return c.value
}

func (c UpdateDataServiceInstanceObsLogRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDataServiceInstanceObsLogRequestDlmType) UnmarshalJSON(b []byte) error {
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
