package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTransferTaskResponse Response Object
type ShowTransferTaskResponse struct {

	// 该转储任务所属通道名称。
	StreamName *string `json:"stream_name,omitempty"`

	// 转储任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 转储任务状态。  - ERROR：错误。 - STARTING：启动中。 - PAUSED：已停止。 - RUNNING：运行中。 - DELETE：已删除。 - ABNORMAL：异常。
	State *ShowTransferTaskResponseState `json:"state,omitempty"`

	// 转储任务类型。  - OBS：转储到OBS。 - MRS：转储到MRS。 - DLI：转储到DLI。 - CLOUDTABLE：转储到CloudTable。 - DWS：转储到DWS。
	DestinationType *ShowTransferTaskResponseDestinationType `json:"destination_type,omitempty"`

	// 转储任务创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 转储任务最近一次转储时间。
	LastTransferTimestamp *int64 `json:"last_transfer_timestamp,omitempty"`

	// 分区转储详情列表。
	Partitions *[]PartitionResult `json:"partitions,omitempty"`

	ObsDestinationDescription *ObsDestinationDescriptorRequest `json:"obs_destination_description,omitempty"`

	DwsDestinationDescripton *DwsDestinationDescriptorRequest `json:"dws_destination_descripton,omitempty"`

	MrsDestinationDescription *MrsDestinationDescriptorRequest `json:"mrs_destination_description,omitempty"`

	DliDestinationDescription *DliDestinationDescriptorRequest `json:"dli_destination_description,omitempty"`

	CloudtableDestinationDescripton *CloudtableDestinationDescriptorRequest `json:"cloudtable_destination_descripton,omitempty"`
	HttpStatusCode                  int                                     `json:"-"`
}

func (o ShowTransferTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransferTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTransferTaskResponse", string(data)}, " ")
}

type ShowTransferTaskResponseState struct {
	value string
}

type ShowTransferTaskResponseStateEnum struct {
	ERROR    ShowTransferTaskResponseState
	STARTING ShowTransferTaskResponseState
	PAUSED   ShowTransferTaskResponseState
	RUNNING  ShowTransferTaskResponseState
	DELETE   ShowTransferTaskResponseState
	ABNORMAL ShowTransferTaskResponseState
}

func GetShowTransferTaskResponseStateEnum() ShowTransferTaskResponseStateEnum {
	return ShowTransferTaskResponseStateEnum{
		ERROR: ShowTransferTaskResponseState{
			value: "ERROR",
		},
		STARTING: ShowTransferTaskResponseState{
			value: "STARTING",
		},
		PAUSED: ShowTransferTaskResponseState{
			value: "PAUSED",
		},
		RUNNING: ShowTransferTaskResponseState{
			value: "RUNNING",
		},
		DELETE: ShowTransferTaskResponseState{
			value: "DELETE",
		},
		ABNORMAL: ShowTransferTaskResponseState{
			value: "ABNORMAL",
		},
	}
}

func (c ShowTransferTaskResponseState) Value() string {
	return c.value
}

func (c ShowTransferTaskResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTransferTaskResponseState) UnmarshalJSON(b []byte) error {
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

type ShowTransferTaskResponseDestinationType struct {
	value string
}

type ShowTransferTaskResponseDestinationTypeEnum struct {
	OBS        ShowTransferTaskResponseDestinationType
	MRS        ShowTransferTaskResponseDestinationType
	DLI        ShowTransferTaskResponseDestinationType
	CLOUDTABLE ShowTransferTaskResponseDestinationType
	DWS        ShowTransferTaskResponseDestinationType
}

func GetShowTransferTaskResponseDestinationTypeEnum() ShowTransferTaskResponseDestinationTypeEnum {
	return ShowTransferTaskResponseDestinationTypeEnum{
		OBS: ShowTransferTaskResponseDestinationType{
			value: "OBS",
		},
		MRS: ShowTransferTaskResponseDestinationType{
			value: "MRS",
		},
		DLI: ShowTransferTaskResponseDestinationType{
			value: "DLI",
		},
		CLOUDTABLE: ShowTransferTaskResponseDestinationType{
			value: "CLOUDTABLE",
		},
		DWS: ShowTransferTaskResponseDestinationType{
			value: "DWS",
		},
	}
}

func (c ShowTransferTaskResponseDestinationType) Value() string {
	return c.value
}

func (c ShowTransferTaskResponseDestinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTransferTaskResponseDestinationType) UnmarshalJSON(b []byte) error {
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
