package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CatalogMetaDataEventRequest 元数据实时同步的事件请求体
type CatalogMetaDataEventRequest struct {

	// 引擎服务名称,DLI DWS MRS
	Engine *CatalogMetaDataEventRequestEngine `json:"engine,omitempty"`

	// 引擎版本信息
	EngineVersion *string `json:"engine_version,omitempty"`

	// 引擎的实例ID, MRS DWS必填
	InstanceId *string `json:"instance_id,omitempty"`

	// 项目ID，DLI必填
	ProjectId *string `json:"project_id,omitempty"`

	// 资产信息
	Events *[]CatalogMetaDataEventInfo `json:"events,omitempty"`
}

func (o CatalogMetaDataEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogMetaDataEventRequest struct{}"
	}

	return strings.Join([]string{"CatalogMetaDataEventRequest", string(data)}, " ")
}

type CatalogMetaDataEventRequestEngine struct {
	value string
}

type CatalogMetaDataEventRequestEngineEnum struct {
	DLI CatalogMetaDataEventRequestEngine
	MRS CatalogMetaDataEventRequestEngine
	DWS CatalogMetaDataEventRequestEngine
}

func GetCatalogMetaDataEventRequestEngineEnum() CatalogMetaDataEventRequestEngineEnum {
	return CatalogMetaDataEventRequestEngineEnum{
		DLI: CatalogMetaDataEventRequestEngine{
			value: "DLI",
		},
		MRS: CatalogMetaDataEventRequestEngine{
			value: "MRS",
		},
		DWS: CatalogMetaDataEventRequestEngine{
			value: "DWS",
		},
	}
}

func (c CatalogMetaDataEventRequestEngine) Value() string {
	return c.value
}

func (c CatalogMetaDataEventRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogMetaDataEventRequestEngine) UnmarshalJSON(b []byte) error {
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
