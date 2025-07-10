package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBackupTransfersRequest Request Object
type ListBackupTransfersRequest struct {

	// 查询记录数
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段
	OrderField *string `json:"order_field,omitempty"`

	// 排序规则
	OrderRule *string `json:"order_rule,omitempty"`

	// 筛选字段
	FilterField *string `json:"filter_field,omitempty"`

	// 筛选关键字
	FilterContent *string `json:"filter_content,omitempty"`

	// 开始时间戳
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 结束时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 转储任务类型，默认为manual manual:手动转储任务 auto:自动转储任务
	TransferType *ListBackupTransfersRequestTransferType `json:"transfer_type,omitempty"`

	// 分页页码
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListBackupTransfersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupTransfersRequest struct{}"
	}

	return strings.Join([]string{"ListBackupTransfersRequest", string(data)}, " ")
}

type ListBackupTransfersRequestTransferType struct {
	value string
}

type ListBackupTransfersRequestTransferTypeEnum struct {
	MANUAL ListBackupTransfersRequestTransferType
	AUTO   ListBackupTransfersRequestTransferType
}

func GetListBackupTransfersRequestTransferTypeEnum() ListBackupTransfersRequestTransferTypeEnum {
	return ListBackupTransfersRequestTransferTypeEnum{
		MANUAL: ListBackupTransfersRequestTransferType{
			value: "manual",
		},
		AUTO: ListBackupTransfersRequestTransferType{
			value: "auto",
		},
	}
}

func (c ListBackupTransfersRequestTransferType) Value() string {
	return c.value
}

func (c ListBackupTransfersRequestTransferType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupTransfersRequestTransferType) UnmarshalJSON(b []byte) error {
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
