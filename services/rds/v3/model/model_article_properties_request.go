package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ArticlePropertiesRequest 项目属性。
type ArticlePropertiesRequest struct {

	// 目标对象名称。
	DestinationObjectName *string `json:"destination_object_name,omitempty"`

	// 目标对象命名空间。
	DestinationObjectOwner *string `json:"destination_object_owner,omitempty"`

	// 插入交付格式。
	InsertDeliveryFormat *ArticlePropertiesRequestInsertDeliveryFormat `json:"insert_delivery_format,omitempty"`

	// 插入存储过程。插入交付格式填call_procedure时该项必填。
	InsertStoredProcedure *string `json:"insert_stored_procedure,omitempty"`

	// 更新交付格式。
	UpdateDeliveryFormat *ArticlePropertiesRequestUpdateDeliveryFormat `json:"update_delivery_format,omitempty"`

	// 更新存储过程。更新交付格式填(m/x/s)call_procedure时该项必填。
	UpdateStoredProcedure *string `json:"update_stored_procedure,omitempty"`

	// 删除交付格式。
	DeleteDeliveryFormat *ArticlePropertiesRequestDeleteDeliveryFormat `json:"delete_delivery_format,omitempty"`

	// 删除存储过程。删除交付格式填(x)call_procedure时该项必填。
	DeleteStoredProcedure *string `json:"delete_stored_procedure,omitempty"`
}

func (o ArticlePropertiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArticlePropertiesRequest struct{}"
	}

	return strings.Join([]string{"ArticlePropertiesRequest", string(data)}, " ")
}

type ArticlePropertiesRequestInsertDeliveryFormat struct {
	value string
}

type ArticlePropertiesRequestInsertDeliveryFormatEnum struct {
	DO_NOT_INSERT              ArticlePropertiesRequestInsertDeliveryFormat
	INSERT                     ArticlePropertiesRequestInsertDeliveryFormat
	INSERT_WITHOUT_COLUMN_LIST ArticlePropertiesRequestInsertDeliveryFormat
	CALL_PROCEDURE             ArticlePropertiesRequestInsertDeliveryFormat
}

func GetArticlePropertiesRequestInsertDeliveryFormatEnum() ArticlePropertiesRequestInsertDeliveryFormatEnum {
	return ArticlePropertiesRequestInsertDeliveryFormatEnum{
		DO_NOT_INSERT: ArticlePropertiesRequestInsertDeliveryFormat{
			value: "do_not_insert",
		},
		INSERT: ArticlePropertiesRequestInsertDeliveryFormat{
			value: "insert",
		},
		INSERT_WITHOUT_COLUMN_LIST: ArticlePropertiesRequestInsertDeliveryFormat{
			value: "insert_without_column_list",
		},
		CALL_PROCEDURE: ArticlePropertiesRequestInsertDeliveryFormat{
			value: "call_procedure",
		},
	}
}

func (c ArticlePropertiesRequestInsertDeliveryFormat) Value() string {
	return c.value
}

func (c ArticlePropertiesRequestInsertDeliveryFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ArticlePropertiesRequestInsertDeliveryFormat) UnmarshalJSON(b []byte) error {
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

type ArticlePropertiesRequestUpdateDeliveryFormat struct {
	value string
}

type ArticlePropertiesRequestUpdateDeliveryFormatEnum struct {
	DO_NOT_UPDATE   ArticlePropertiesRequestUpdateDeliveryFormat
	UPDATE          ArticlePropertiesRequestUpdateDeliveryFormat
	CALL_PROCEDURE  ArticlePropertiesRequestUpdateDeliveryFormat
	MCALL_PROCEDURE ArticlePropertiesRequestUpdateDeliveryFormat
	XCALL_PROCEDURE ArticlePropertiesRequestUpdateDeliveryFormat
	SCALL_PROCEDURE ArticlePropertiesRequestUpdateDeliveryFormat
}

func GetArticlePropertiesRequestUpdateDeliveryFormatEnum() ArticlePropertiesRequestUpdateDeliveryFormatEnum {
	return ArticlePropertiesRequestUpdateDeliveryFormatEnum{
		DO_NOT_UPDATE: ArticlePropertiesRequestUpdateDeliveryFormat{
			value: "do_not_update",
		},
		UPDATE: ArticlePropertiesRequestUpdateDeliveryFormat{
			value: "update",
		},
		CALL_PROCEDURE: ArticlePropertiesRequestUpdateDeliveryFormat{
			value: "call_procedure",
		},
		MCALL_PROCEDURE: ArticlePropertiesRequestUpdateDeliveryFormat{
			value: "mcall_procedure",
		},
		XCALL_PROCEDURE: ArticlePropertiesRequestUpdateDeliveryFormat{
			value: "xcall_procedure",
		},
		SCALL_PROCEDURE: ArticlePropertiesRequestUpdateDeliveryFormat{
			value: "scall_procedure",
		},
	}
}

func (c ArticlePropertiesRequestUpdateDeliveryFormat) Value() string {
	return c.value
}

func (c ArticlePropertiesRequestUpdateDeliveryFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ArticlePropertiesRequestUpdateDeliveryFormat) UnmarshalJSON(b []byte) error {
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

type ArticlePropertiesRequestDeleteDeliveryFormat struct {
	value string
}

type ArticlePropertiesRequestDeleteDeliveryFormatEnum struct {
	DO_NOT_DELETE   ArticlePropertiesRequestDeleteDeliveryFormat
	DELETE          ArticlePropertiesRequestDeleteDeliveryFormat
	CALL_PROCEDURE  ArticlePropertiesRequestDeleteDeliveryFormat
	XCALL_PROCEDURE ArticlePropertiesRequestDeleteDeliveryFormat
}

func GetArticlePropertiesRequestDeleteDeliveryFormatEnum() ArticlePropertiesRequestDeleteDeliveryFormatEnum {
	return ArticlePropertiesRequestDeleteDeliveryFormatEnum{
		DO_NOT_DELETE: ArticlePropertiesRequestDeleteDeliveryFormat{
			value: "do_not_delete",
		},
		DELETE: ArticlePropertiesRequestDeleteDeliveryFormat{
			value: "delete",
		},
		CALL_PROCEDURE: ArticlePropertiesRequestDeleteDeliveryFormat{
			value: "call_procedure",
		},
		XCALL_PROCEDURE: ArticlePropertiesRequestDeleteDeliveryFormat{
			value: "xcall_procedure",
		},
	}
}

func (c ArticlePropertiesRequestDeleteDeliveryFormat) Value() string {
	return c.value
}

func (c ArticlePropertiesRequestDeleteDeliveryFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ArticlePropertiesRequestDeleteDeliveryFormat) UnmarshalJSON(b []byte) error {
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
