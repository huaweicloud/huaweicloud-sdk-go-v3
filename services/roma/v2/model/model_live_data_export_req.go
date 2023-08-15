package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LiveDataExportReq struct {

	// API所属的应用ID
	AppId *string `json:"app_id,omitempty"`

	// 导出的API定义的格式
	Format *LiveDataExportReqFormat `json:"format,omitempty"`

	// 导出的自定义后端API ID列表
	Apis *[]string `json:"apis,omitempty"`

	// 导出的后端API状态： - 1：待开发 - 3：开发中 - 4：已部署
	Status *int32 `json:"status,omitempty"`
}

func (o LiveDataExportReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveDataExportReq struct{}"
	}

	return strings.Join([]string{"LiveDataExportReq", string(data)}, " ")
}

type LiveDataExportReqFormat struct {
	value string
}

type LiveDataExportReqFormatEnum struct {
	JSON LiveDataExportReqFormat
	YAML LiveDataExportReqFormat
	YML  LiveDataExportReqFormat
}

func GetLiveDataExportReqFormatEnum() LiveDataExportReqFormatEnum {
	return LiveDataExportReqFormatEnum{
		JSON: LiveDataExportReqFormat{
			value: "json",
		},
		YAML: LiveDataExportReqFormat{
			value: "yaml",
		},
		YML: LiveDataExportReqFormat{
			value: "yml",
		},
	}
}

func (c LiveDataExportReqFormat) Value() string {
	return c.value
}

func (c LiveDataExportReqFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveDataExportReqFormat) UnmarshalJSON(b []byte) error {
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
