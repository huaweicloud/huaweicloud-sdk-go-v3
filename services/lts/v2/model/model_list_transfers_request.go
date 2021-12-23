package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTransfersRequest struct {
	// 日志转储类型。OBS指OBS日志转储，DIS指DIS日志转储，DMS指DMS日志转储

	LogTransferType *ListTransfersRequestLogTransferType `json:"log_transfer_type,omitempty"`
	// 日志组名称

	LogGroupName *string `json:"log_group_name,omitempty"`
	// 日志流名称

	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o ListTransfersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransfersRequest struct{}"
	}

	return strings.Join([]string{"ListTransfersRequest", string(data)}, " ")
}

type ListTransfersRequestLogTransferType struct {
	value string
}

type ListTransfersRequestLogTransferTypeEnum struct {
	OBS ListTransfersRequestLogTransferType
	DIS ListTransfersRequestLogTransferType
	DMS ListTransfersRequestLogTransferType
}

func GetListTransfersRequestLogTransferTypeEnum() ListTransfersRequestLogTransferTypeEnum {
	return ListTransfersRequestLogTransferTypeEnum{
		OBS: ListTransfersRequestLogTransferType{
			value: "OBS",
		},
		DIS: ListTransfersRequestLogTransferType{
			value: "DIS",
		},
		DMS: ListTransfersRequestLogTransferType{
			value: "DMS",
		},
	}
}

func (c ListTransfersRequestLogTransferType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTransfersRequestLogTransferType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
