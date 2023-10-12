package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAllCatalogListRequest Request Object
type ListAllCatalogListRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ListAllCatalogListRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// 查询起始坐标, 即跳过前X条数据
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAllCatalogListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllCatalogListRequest struct{}"
	}

	return strings.Join([]string{"ListAllCatalogListRequest", string(data)}, " ")
}

type ListAllCatalogListRequestDlmType struct {
	value string
}

type ListAllCatalogListRequestDlmTypeEnum struct {
	SHARED    ListAllCatalogListRequestDlmType
	EXCLUSIVE ListAllCatalogListRequestDlmType
}

func GetListAllCatalogListRequestDlmTypeEnum() ListAllCatalogListRequestDlmTypeEnum {
	return ListAllCatalogListRequestDlmTypeEnum{
		SHARED: ListAllCatalogListRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListAllCatalogListRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListAllCatalogListRequestDlmType) Value() string {
	return c.value
}

func (c ListAllCatalogListRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllCatalogListRequestDlmType) UnmarshalJSON(b []byte) error {
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
