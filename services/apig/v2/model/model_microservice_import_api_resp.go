package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroserviceImportApiResp 导入的微服务API的响应对象
type MicroserviceImportApiResp struct {

	// API名称
	Name *string `json:"name,omitempty"`

	// API请求路径
	ReqUri *string `json:"req_uri,omitempty"`

	// API请求方法
	ReqMethod *string `json:"req_method,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：SWA
	MatchMode *MicroserviceImportApiRespMatchMode `json:"match_mode,omitempty"`
}

func (o MicroserviceImportApiResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroserviceImportApiResp struct{}"
	}

	return strings.Join([]string{"MicroserviceImportApiResp", string(data)}, " ")
}

type MicroserviceImportApiRespMatchMode struct {
	value string
}

type MicroserviceImportApiRespMatchModeEnum struct {
	SWA    MicroserviceImportApiRespMatchMode
	NORMAL MicroserviceImportApiRespMatchMode
}

func GetMicroserviceImportApiRespMatchModeEnum() MicroserviceImportApiRespMatchModeEnum {
	return MicroserviceImportApiRespMatchModeEnum{
		SWA: MicroserviceImportApiRespMatchMode{
			value: "SWA",
		},
		NORMAL: MicroserviceImportApiRespMatchMode{
			value: "NORMAL",
		},
	}
}

func (c MicroserviceImportApiRespMatchMode) Value() string {
	return c.value
}

func (c MicroserviceImportApiRespMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroserviceImportApiRespMatchMode) UnmarshalJSON(b []byte) error {
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
