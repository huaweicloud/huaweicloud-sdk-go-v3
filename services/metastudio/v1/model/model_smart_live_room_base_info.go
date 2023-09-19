package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SmartLiveRoomBaseInfo struct {

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播间名称
	RoomName *string `json:"room_name,omitempty"`

	// 直播间描述。
	RoomDescription *string `json:"room_description,omitempty"`

	// 直播间封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	// 数字人模型信息
	ModelInfos *[]ModelInfo `json:"model_infos,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	// 开始直播时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	LastJobStartTime *string `json:"last_job_start_time,omitempty"`

	// 结束直播时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	LastJobEndTime *string `json:"last_job_end_time,omitempty"`

	// 当前直播状态 - WAITING：任务等待执行 - PROCESSING：任务执行中 - SUCCEED：任务处理成功 - FAILED：任务处理时变 - CANCELED：任务取消
	LastJobStatus *SmartLiveRoomBaseInfoLastJobStatus `json:"last_job_status,omitempty"`
}

func (o SmartLiveRoomBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartLiveRoomBaseInfo struct{}"
	}

	return strings.Join([]string{"SmartLiveRoomBaseInfo", string(data)}, " ")
}

type SmartLiveRoomBaseInfoLastJobStatus struct {
	value string
}

type SmartLiveRoomBaseInfoLastJobStatusEnum struct {
	WAITING    SmartLiveRoomBaseInfoLastJobStatus
	PROCESSING SmartLiveRoomBaseInfoLastJobStatus
	SUCCEED    SmartLiveRoomBaseInfoLastJobStatus
	FAILED     SmartLiveRoomBaseInfoLastJobStatus
	CANCELED   SmartLiveRoomBaseInfoLastJobStatus
}

func GetSmartLiveRoomBaseInfoLastJobStatusEnum() SmartLiveRoomBaseInfoLastJobStatusEnum {
	return SmartLiveRoomBaseInfoLastJobStatusEnum{
		WAITING: SmartLiveRoomBaseInfoLastJobStatus{
			value: "WAITING",
		},
		PROCESSING: SmartLiveRoomBaseInfoLastJobStatus{
			value: "PROCESSING",
		},
		SUCCEED: SmartLiveRoomBaseInfoLastJobStatus{
			value: "SUCCEED",
		},
		FAILED: SmartLiveRoomBaseInfoLastJobStatus{
			value: "FAILED",
		},
		CANCELED: SmartLiveRoomBaseInfoLastJobStatus{
			value: "CANCELED",
		},
	}
}

func (c SmartLiveRoomBaseInfoLastJobStatus) Value() string {
	return c.value
}

func (c SmartLiveRoomBaseInfoLastJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLiveRoomBaseInfoLastJobStatus) UnmarshalJSON(b []byte) error {
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
