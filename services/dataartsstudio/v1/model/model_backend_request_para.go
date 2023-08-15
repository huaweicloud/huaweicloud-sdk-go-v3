package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackendRequestPara struct {

	// api请求参数名称
	Name *string `json:"name,omitempty"`

	// 参数位置
	Position *BackendRequestParaPosition `json:"position,omitempty"`

	// 对应的后端参数
	BackendParaName *string `json:"backend_para_name,omitempty"`
}

func (o BackendRequestPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendRequestPara struct{}"
	}

	return strings.Join([]string{"BackendRequestPara", string(data)}, " ")
}

type BackendRequestParaPosition struct {
	value string
}

type BackendRequestParaPositionEnum struct {
	REQUEST_PARAMETER_POSITION_PATH   BackendRequestParaPosition
	REQUEST_PARAMETER_POSITION_HEADER BackendRequestParaPosition
	REQUEST_PARAMETER_POSITION_QUERY  BackendRequestParaPosition
}

func GetBackendRequestParaPositionEnum() BackendRequestParaPositionEnum {
	return BackendRequestParaPositionEnum{
		REQUEST_PARAMETER_POSITION_PATH: BackendRequestParaPosition{
			value: "REQUEST_PARAMETER_POSITION_PATH",
		},
		REQUEST_PARAMETER_POSITION_HEADER: BackendRequestParaPosition{
			value: "REQUEST_PARAMETER_POSITION_HEADER",
		},
		REQUEST_PARAMETER_POSITION_QUERY: BackendRequestParaPosition{
			value: "REQUEST_PARAMETER_POSITION_QUERY",
		},
	}
}

func (c BackendRequestParaPosition) Value() string {
	return c.value
}

func (c BackendRequestParaPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendRequestParaPosition) UnmarshalJSON(b []byte) error {
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
