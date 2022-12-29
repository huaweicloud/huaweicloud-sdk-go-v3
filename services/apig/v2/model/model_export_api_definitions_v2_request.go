package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ExportApiDefinitionsV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// OpenAPI版本
	OasVersion *ExportApiDefinitionsV2RequestOasVersion `json:"oas_version,omitempty"`

	Body *ExportOpenApiReq `json:"body,omitempty"`
}

func (o ExportApiDefinitionsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ExportApiDefinitionsV2Request", string(data)}, " ")
}

type ExportApiDefinitionsV2RequestOasVersion struct {
	value string
}

type ExportApiDefinitionsV2RequestOasVersionEnum struct {
	E_2_0 ExportApiDefinitionsV2RequestOasVersion
	E_3_0 ExportApiDefinitionsV2RequestOasVersion
}

func GetExportApiDefinitionsV2RequestOasVersionEnum() ExportApiDefinitionsV2RequestOasVersionEnum {
	return ExportApiDefinitionsV2RequestOasVersionEnum{
		E_2_0: ExportApiDefinitionsV2RequestOasVersion{
			value: "2.0",
		},
		E_3_0: ExportApiDefinitionsV2RequestOasVersion{
			value: "3.0",
		},
	}
}

func (c ExportApiDefinitionsV2RequestOasVersion) Value() string {
	return c.value
}

func (c ExportApiDefinitionsV2RequestOasVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportApiDefinitionsV2RequestOasVersion) UnmarshalJSON(b []byte) error {
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
