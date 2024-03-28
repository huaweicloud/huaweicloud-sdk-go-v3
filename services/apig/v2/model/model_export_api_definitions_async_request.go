package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportApiDefinitionsAsyncRequest Request Object
type ExportApiDefinitionsAsyncRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// OpenAPI版本
	OasVersion *ExportApiDefinitionsAsyncRequestOasVersion `json:"oas_version,omitempty"`

	Body *ExportOpenApiReq `json:"body,omitempty"`
}

func (o ExportApiDefinitionsAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportApiDefinitionsAsyncRequest struct{}"
	}

	return strings.Join([]string{"ExportApiDefinitionsAsyncRequest", string(data)}, " ")
}

type ExportApiDefinitionsAsyncRequestOasVersion struct {
	value string
}

type ExportApiDefinitionsAsyncRequestOasVersionEnum struct {
	E_2_0 ExportApiDefinitionsAsyncRequestOasVersion
	E_3_0 ExportApiDefinitionsAsyncRequestOasVersion
}

func GetExportApiDefinitionsAsyncRequestOasVersionEnum() ExportApiDefinitionsAsyncRequestOasVersionEnum {
	return ExportApiDefinitionsAsyncRequestOasVersionEnum{
		E_2_0: ExportApiDefinitionsAsyncRequestOasVersion{
			value: "2.0",
		},
		E_3_0: ExportApiDefinitionsAsyncRequestOasVersion{
			value: "3.0",
		},
	}
}

func (c ExportApiDefinitionsAsyncRequestOasVersion) Value() string {
	return c.value
}

func (c ExportApiDefinitionsAsyncRequestOasVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportApiDefinitionsAsyncRequestOasVersion) UnmarshalJSON(b []byte) error {
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
