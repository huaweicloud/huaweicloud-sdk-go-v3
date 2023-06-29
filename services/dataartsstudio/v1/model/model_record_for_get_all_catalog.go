package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecordForGetAllCatalog struct {

	// 目录编号
	CatalogId *string `json:"catalog_id,omitempty"`

	// 父目录编号
	Pid *string `json:"pid,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 目录描述
	Description *string `json:"description,omitempty"`

	// 目录类型
	ApiCatalogType *RecordForGetAllCatalogApiCatalogType `json:"api_catalog_type,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 更新者
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o RecordForGetAllCatalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordForGetAllCatalog struct{}"
	}

	return strings.Join([]string{"RecordForGetAllCatalog", string(data)}, " ")
}

type RecordForGetAllCatalogApiCatalogType struct {
	value string
}

type RecordForGetAllCatalogApiCatalogTypeEnum struct {
	CATALOG RecordForGetAllCatalogApiCatalogType
	API     RecordForGetAllCatalogApiCatalogType
}

func GetRecordForGetAllCatalogApiCatalogTypeEnum() RecordForGetAllCatalogApiCatalogTypeEnum {
	return RecordForGetAllCatalogApiCatalogTypeEnum{
		CATALOG: RecordForGetAllCatalogApiCatalogType{
			value: "CATALOG",
		},
		API: RecordForGetAllCatalogApiCatalogType{
			value: "API",
		},
	}
}

func (c RecordForGetAllCatalogApiCatalogType) Value() string {
	return c.value
}

func (c RecordForGetAllCatalogApiCatalogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordForGetAllCatalogApiCatalogType) UnmarshalJSON(b []byte) error {
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
