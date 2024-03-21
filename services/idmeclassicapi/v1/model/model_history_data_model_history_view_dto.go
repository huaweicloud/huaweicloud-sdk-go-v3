package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HistoryDataModelHistoryViewDto struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 修改时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 系统版本，用于存储MONGO。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// 操作类型，用于存储MONGO。
	RdmOperationType *HistoryDataModelHistoryViewDtoRdmOperationType `json:"rdmOperationType,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 删除标志。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	Tenant *TenantHistoryViewDto `json:"tenant,omitempty"`

	// 类名称。
	ClassName *string `json:"className,omitempty"`
}

func (o HistoryDataModelHistoryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryDataModelHistoryViewDto struct{}"
	}

	return strings.Join([]string{"HistoryDataModelHistoryViewDto", string(data)}, " ")
}

type HistoryDataModelHistoryViewDtoRdmOperationType struct {
	value string
}

type HistoryDataModelHistoryViewDtoRdmOperationTypeEnum struct {
	CREATE        HistoryDataModelHistoryViewDtoRdmOperationType
	UPDATE        HistoryDataModelHistoryViewDtoRdmOperationType
	LOGICALDELETE HistoryDataModelHistoryViewDtoRdmOperationType
	DELETE        HistoryDataModelHistoryViewDtoRdmOperationType
	CASCADE       HistoryDataModelHistoryViewDtoRdmOperationType
}

func GetHistoryDataModelHistoryViewDtoRdmOperationTypeEnum() HistoryDataModelHistoryViewDtoRdmOperationTypeEnum {
	return HistoryDataModelHistoryViewDtoRdmOperationTypeEnum{
		CREATE: HistoryDataModelHistoryViewDtoRdmOperationType{
			value: "CREATE",
		},
		UPDATE: HistoryDataModelHistoryViewDtoRdmOperationType{
			value: "UPDATE",
		},
		LOGICALDELETE: HistoryDataModelHistoryViewDtoRdmOperationType{
			value: "LOGICALDELETE",
		},
		DELETE: HistoryDataModelHistoryViewDtoRdmOperationType{
			value: "DELETE",
		},
		CASCADE: HistoryDataModelHistoryViewDtoRdmOperationType{
			value: "CASCADE",
		},
	}
}

func (c HistoryDataModelHistoryViewDtoRdmOperationType) Value() string {
	return c.value
}

func (c HistoryDataModelHistoryViewDtoRdmOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HistoryDataModelHistoryViewDtoRdmOperationType) UnmarshalJSON(b []byte) error {
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
