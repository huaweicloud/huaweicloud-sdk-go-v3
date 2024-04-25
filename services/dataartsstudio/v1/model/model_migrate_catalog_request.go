package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MigrateCatalogRequest Request Object
type MigrateCatalogRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *MigrateCatalogRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 目录编号。
	CatalogId string `json:"catalog_id"`

	Body *CatalogMoveParaDto `json:"body,omitempty"`
}

func (o MigrateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateCatalogRequest struct{}"
	}

	return strings.Join([]string{"MigrateCatalogRequest", string(data)}, " ")
}

type MigrateCatalogRequestDlmType struct {
	value string
}

type MigrateCatalogRequestDlmTypeEnum struct {
	SHARED    MigrateCatalogRequestDlmType
	EXCLUSIVE MigrateCatalogRequestDlmType
}

func GetMigrateCatalogRequestDlmTypeEnum() MigrateCatalogRequestDlmTypeEnum {
	return MigrateCatalogRequestDlmTypeEnum{
		SHARED: MigrateCatalogRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: MigrateCatalogRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c MigrateCatalogRequestDlmType) Value() string {
	return c.value
}

func (c MigrateCatalogRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrateCatalogRequestDlmType) UnmarshalJSON(b []byte) error {
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
