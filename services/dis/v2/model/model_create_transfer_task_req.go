package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTransferTaskReq struct {

	// 转储任务类型。  - OBS：转储到OBS - MRS：转储到MRS - DLI：转储到DLI - CLOUDTABLE：转储到CloudTable - DWS：转储到DWS
	DestinationType CreateTransferTaskReqDestinationType `json:"destination_type"`

	ObsDestinationDescriptor *ObsDestinationDescriptorRequest `json:"obs_destination_descriptor,omitempty"`
}

func (o CreateTransferTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferTaskReq struct{}"
	}

	return strings.Join([]string{"CreateTransferTaskReq", string(data)}, " ")
}

type CreateTransferTaskReqDestinationType struct {
	value string
}

type CreateTransferTaskReqDestinationTypeEnum struct {
	OBS CreateTransferTaskReqDestinationType
}

func GetCreateTransferTaskReqDestinationTypeEnum() CreateTransferTaskReqDestinationTypeEnum {
	return CreateTransferTaskReqDestinationTypeEnum{
		OBS: CreateTransferTaskReqDestinationType{
			value: "OBS",
		},
	}
}

func (c CreateTransferTaskReqDestinationType) Value() string {
	return c.value
}

func (c CreateTransferTaskReqDestinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferTaskReqDestinationType) UnmarshalJSON(b []byte) error {
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
