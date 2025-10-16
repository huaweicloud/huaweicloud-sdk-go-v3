package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ArticlePropertiesResponse 项目属性。
type ArticlePropertiesResponse struct {

	// 目标对象名称。
	DestinationObjectName *string `json:"destination_object_name,omitempty"`

	// 目标对象命名空间。
	DestinationObjectOwner *string `json:"destination_object_owner,omitempty"`

	// 插入交付格式。
	InsertDeliveryFormat *ArticlePropertiesResponseInsertDeliveryFormat `json:"insert_delivery_format,omitempty"`

	// 插入存储过程。插入交付格式填call_procedure时该项必填。
	InsertStoredProcedure *string `json:"insert_stored_procedure,omitempty"`

	// 更新交付格式。
	UpdateDeliveryFormat *ArticlePropertiesResponseUpdateDeliveryFormat `json:"update_delivery_format,omitempty"`

	// 更新存储过程。更新交付格式填(m/x/s)call_procedure时该项必填。
	UpdateStoredProcedure *string `json:"update_stored_procedure,omitempty"`

	// 删除交付格式。
	DeleteDeliveryFormat *ArticlePropertiesResponseDeleteDeliveryFormat `json:"delete_delivery_format,omitempty"`

	// 删除存储过程。删除交付格式填(x)call_procedure时该项必填。
	DeleteStoredProcedure *string `json:"delete_stored_procedure,omitempty"`
}

func (o ArticlePropertiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArticlePropertiesResponse struct{}"
	}

	return strings.Join([]string{"ArticlePropertiesResponse", string(data)}, " ")
}

type ArticlePropertiesResponseInsertDeliveryFormat struct {
	value string
}

type ArticlePropertiesResponseInsertDeliveryFormatEnum struct {
	DO_NOT_INSERT              ArticlePropertiesResponseInsertDeliveryFormat
	INSERT                     ArticlePropertiesResponseInsertDeliveryFormat
	INSERT_WITHOUT_COLUMN_LIST ArticlePropertiesResponseInsertDeliveryFormat
	CALL_PROCEDURE             ArticlePropertiesResponseInsertDeliveryFormat
}

func GetArticlePropertiesResponseInsertDeliveryFormatEnum() ArticlePropertiesResponseInsertDeliveryFormatEnum {
	return ArticlePropertiesResponseInsertDeliveryFormatEnum{
		DO_NOT_INSERT: ArticlePropertiesResponseInsertDeliveryFormat{
			value: "do_not_insert",
		},
		INSERT: ArticlePropertiesResponseInsertDeliveryFormat{
			value: "insert",
		},
		INSERT_WITHOUT_COLUMN_LIST: ArticlePropertiesResponseInsertDeliveryFormat{
			value: "insert_without_column_list",
		},
		CALL_PROCEDURE: ArticlePropertiesResponseInsertDeliveryFormat{
			value: "call_procedure",
		},
	}
}

func (c ArticlePropertiesResponseInsertDeliveryFormat) Value() string {
	return c.value
}

func (c ArticlePropertiesResponseInsertDeliveryFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ArticlePropertiesResponseInsertDeliveryFormat) UnmarshalJSON(b []byte) error {
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

type ArticlePropertiesResponseUpdateDeliveryFormat struct {
	value string
}

type ArticlePropertiesResponseUpdateDeliveryFormatEnum struct {
	DO_NOT_UPDATE   ArticlePropertiesResponseUpdateDeliveryFormat
	UPDATE          ArticlePropertiesResponseUpdateDeliveryFormat
	CALL_PROCEDURE  ArticlePropertiesResponseUpdateDeliveryFormat
	MCALL_PROCEDURE ArticlePropertiesResponseUpdateDeliveryFormat
	XCALL_PROCEDURE ArticlePropertiesResponseUpdateDeliveryFormat
	SCALL_PROCEDURE ArticlePropertiesResponseUpdateDeliveryFormat
}

func GetArticlePropertiesResponseUpdateDeliveryFormatEnum() ArticlePropertiesResponseUpdateDeliveryFormatEnum {
	return ArticlePropertiesResponseUpdateDeliveryFormatEnum{
		DO_NOT_UPDATE: ArticlePropertiesResponseUpdateDeliveryFormat{
			value: "do_not_update",
		},
		UPDATE: ArticlePropertiesResponseUpdateDeliveryFormat{
			value: "update",
		},
		CALL_PROCEDURE: ArticlePropertiesResponseUpdateDeliveryFormat{
			value: "call_procedure",
		},
		MCALL_PROCEDURE: ArticlePropertiesResponseUpdateDeliveryFormat{
			value: "mcall_procedure",
		},
		XCALL_PROCEDURE: ArticlePropertiesResponseUpdateDeliveryFormat{
			value: "xcall_procedure",
		},
		SCALL_PROCEDURE: ArticlePropertiesResponseUpdateDeliveryFormat{
			value: "scall_procedure",
		},
	}
}

func (c ArticlePropertiesResponseUpdateDeliveryFormat) Value() string {
	return c.value
}

func (c ArticlePropertiesResponseUpdateDeliveryFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ArticlePropertiesResponseUpdateDeliveryFormat) UnmarshalJSON(b []byte) error {
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

type ArticlePropertiesResponseDeleteDeliveryFormat struct {
	value string
}

type ArticlePropertiesResponseDeleteDeliveryFormatEnum struct {
	DO_NOT_DELETE   ArticlePropertiesResponseDeleteDeliveryFormat
	DELETE          ArticlePropertiesResponseDeleteDeliveryFormat
	CALL_PROCEDURE  ArticlePropertiesResponseDeleteDeliveryFormat
	XCALL_PROCEDURE ArticlePropertiesResponseDeleteDeliveryFormat
}

func GetArticlePropertiesResponseDeleteDeliveryFormatEnum() ArticlePropertiesResponseDeleteDeliveryFormatEnum {
	return ArticlePropertiesResponseDeleteDeliveryFormatEnum{
		DO_NOT_DELETE: ArticlePropertiesResponseDeleteDeliveryFormat{
			value: "do_not_delete",
		},
		DELETE: ArticlePropertiesResponseDeleteDeliveryFormat{
			value: "delete",
		},
		CALL_PROCEDURE: ArticlePropertiesResponseDeleteDeliveryFormat{
			value: "call_procedure",
		},
		XCALL_PROCEDURE: ArticlePropertiesResponseDeleteDeliveryFormat{
			value: "xcall_procedure",
		},
	}
}

func (c ArticlePropertiesResponseDeleteDeliveryFormat) Value() string {
	return c.value
}

func (c ArticlePropertiesResponseDeleteDeliveryFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ArticlePropertiesResponseDeleteDeliveryFormat) UnmarshalJSON(b []byte) error {
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
