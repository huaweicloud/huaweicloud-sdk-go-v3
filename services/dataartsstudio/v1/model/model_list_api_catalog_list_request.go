package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApiCatalogListRequest Request Object
type ListApiCatalogListRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ListApiCatalogListRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// 查询起始坐标, 即跳过前X条数据
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
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
