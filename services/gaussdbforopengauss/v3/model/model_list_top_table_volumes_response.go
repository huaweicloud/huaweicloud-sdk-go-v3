package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTopTableVolumesResponse Response Object
type ListTopTableVolumesResponse struct {

	// **参数解释**: 数据库表占用空间列表。
	TableVolumes *[]TableVolumeResult `json:"table_volumes,omitempty"`

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 任务状态。 **取值范围**: - RUNNING：运行中。 - ERROR：运行异常。 - FINISHED： 运行结束。
	State          *ListTopTableVolumesResponseState `json:"state,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListTopTableVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopTableVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListTopTableVolumesResponse", string(data)}, " ")
}

type ListTopTableVolumesResponseState struct {
	value string
}

type ListTopTableVolumesResponseStateEnum struct {
	RUNNING  ListTopTableVolumesResponseState
	ERROR    ListTopTableVolumesResponseState
	FINISHED ListTopTableVolumesResponseState
}

func GetListTopTableVolumesResponseStateEnum() ListTopTableVolumesResponseStateEnum {
	return ListTopTableVolumesResponseStateEnum{
		RUNNING: ListTopTableVolumesResponseState{
			value: "RUNNING",
		},
		ERROR: ListTopTableVolumesResponseState{
			value: "ERROR",
		},
		FINISHED: ListTopTableVolumesResponseState{
			value: "FINISHED",
		},
	}
}

func (c ListTopTableVolumesResponseState) Value() string {
	return c.value
}

func (c ListTopTableVolumesResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTopTableVolumesResponseState) UnmarshalJSON(b []byte) error {
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
