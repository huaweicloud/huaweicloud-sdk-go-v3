package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadMetadataRequest Request Object
type DownloadMetadataRequest struct {

	// 认证配置id。
	AuthConfigId string `json:"auth_config_id"`

	// 元数据类型。 - IDP：身份认证提供方 - SP：服务提供方
	Type DownloadMetadataRequestType `json:"type"`
}

func (o DownloadMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadMetadataRequest struct{}"
	}

	return strings.Join([]string{"DownloadMetadataRequest", string(data)}, " ")
}

type DownloadMetadataRequestType struct {
	value string
}

type DownloadMetadataRequestTypeEnum struct {
	IDP DownloadMetadataRequestType
	SP  DownloadMetadataRequestType
}

func GetDownloadMetadataRequestTypeEnum() DownloadMetadataRequestTypeEnum {
	return DownloadMetadataRequestTypeEnum{
		IDP: DownloadMetadataRequestType{
			value: "IDP",
		},
		SP: DownloadMetadataRequestType{
			value: "SP",
		},
	}
}

func (c DownloadMetadataRequestType) Value() string {
	return c.value
}

func (c DownloadMetadataRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadMetadataRequestType) UnmarshalJSON(b []byte) error {
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
