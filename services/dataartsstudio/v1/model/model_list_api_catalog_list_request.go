package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApiCatalogListRequest Request Object
type ListApiCatalogListRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListApiCatalogListRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 目录编号。
	CatalogId string `json:"catalog_id"`

	// 查询起始坐标, 即跳过前X条数据。
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApiCatalogListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiCatalogListRequest struct{}"
	}

	return strings.Join([]string{"ListApiCatalogListRequest", string(data)}, " ")
}

type ListApiCatalogListRequestDlmType struct {
	value string
}

type ListApiCatalogListRequestDlmTypeEnum struct {
	SHARED    ListApiCatalogListRequestDlmType
	EXCLUSIVE ListApiCatalogListRequestDlmType
}

func GetListApiCatalogListRequestDlmTypeEnum() ListApiCatalogListRequestDlmTypeEnum {
	return ListApiCatalogListRequestDlmTypeEnum{
		SHARED: ListApiCatalogListRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApiCatalogListRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApiCatalogListRequestDlmType) Value() string {
	return c.value
}

func (c ListApiCatalogListRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiCatalogListRequestDlmType) UnmarshalJSON(b []byte) error {
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
