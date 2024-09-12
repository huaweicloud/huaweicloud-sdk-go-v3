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

	// 租户id
	ProjectId *string `json:"project_id,omitempty"`

	// 直播间名称
	RoomName *string `json:"room_name,omitempty"`

	// 直播间类型。 * NORMAL：普通直播间，直播间一直存在，可以反复开播 * TEMP：临时直播间，直播任务结束后自动清理直播间。 * TEMPLATE：直播间模板。
	RoomType *SmartLiveRoomBaseInfoRoomType `json:"room_type,omitempty"`

	// 直播间配置状态。 - ENABLE: 直播间正常可用。 - DISABLE： 直播间不可用。不可用原因在error_info中说明。 - BLOCKED：直播间被冻结。冻结原因在error_info中说明。
	RoomState *SmartLiveRoomBaseInfoRoomState `json:"room_state,omitempty"`

	// 横竖屏类型。默认值为：VERTICAL。 * LANDSCAPE：横屏。 * VERTICAL： 竖屏。
	ViewMode *SmartLiveRoomBaseInfoViewMode `json:"view_mode,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	SharedConfig *SharedConfig `json:"shared_config,omitempty"`

	// 直播间描述。
	RoomDescription *string `json:"room_description,omitempty"`

	// 直播间封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	// 直播间封面图URL
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 数字人模型信息
	ModelInfos *[]ModelInfo `json:"model_infos,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// 开始直播时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	LastJobStartTime *string `json:"last_job_start_time,omitempty"`

	// 结束直播时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	LastJobEndTime *string `json:"last_job_end_time,omitempty"`

	// 当前直播状态 - WAITING：任务等待执行 - PROCESSING：任务执行中 - SUCCEED：任务处理成功 - FAILED：任务处理时变 - CANCELED：任务取消 - BLOCKED：任务被冻结
	LastJobStatus *SmartLiveRoomBaseInfoLastJobStatus `json:"last_job_status,omitempty"`

	// 私有数据，用户填写，原样带回。
	PrivData *string `json:"priv_data,omitempty"`
}

func (o SmartLiveRoomBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartLiveRoomBaseInfo struct{}"
	}

	return strings.Join([]string{"SmartLiveRoomBaseInfo", string(data)}, " ")
}

type SmartLiveRoomBaseInfoRoomType struct {
	value string
}

type SmartLiveRoomBaseInfoRoomTypeEnum struct {
	NORMAL   SmartLiveRoomBaseInfoRoomType
	TEMP     SmartLiveRoomBaseInfoRoomType
	TEMPLATE SmartLiveRoomBaseInfoRoomType
}

func GetSmartLiveRoomBaseInfoRoomTypeEnum() SmartLiveRoomBaseInfoRoomTypeEnum {
	return SmartLiveRoomBaseInfoRoomTypeEnum{
		NORMAL: SmartLiveRoomBaseInfoRoomType{
			value: "NORMAL",
		},
		TEMP: SmartLiveRoomBaseInfoRoomType{
			value: "TEMP",
		},
		TEMPLATE: SmartLiveRoomBaseInfoRoomType{
			value: "TEMPLATE",
		},
	}
}

func (c SmartLiveRoomBaseInfoRoomType) Value() string {
	return c.value
}

func (c SmartLiveRoomBaseInfoRoomType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLiveRoomBaseInfoRoomType) UnmarshalJSON(b []byte) error {
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

type SmartLiveRoomBaseInfoRoomState struct {
	value string
}

type SmartLiveRoomBaseInfoRoomStateEnum struct {
	ENABLE  SmartLiveRoomBaseInfoRoomState
	DISABLE SmartLiveRoomBaseInfoRoomState
	BLOCKED SmartLiveRoomBaseInfoRoomState
}

func GetSmartLiveRoomBaseInfoRoomStateEnum() SmartLiveRoomBaseInfoRoomStateEnum {
	return SmartLiveRoomBaseInfoRoomStateEnum{
		ENABLE: SmartLiveRoomBaseInfoRoomState{
			value: "ENABLE",
		},
		DISABLE: SmartLiveRoomBaseInfoRoomState{
			value: "DISABLE",
		},
		BLOCKED: SmartLiveRoomBaseInfoRoomState{
			value: "BLOCKED",
		},
	}
}

func (c SmartLiveRoomBaseInfoRoomState) Value() string {
	return c.value
}

func (c SmartLiveRoomBaseInfoRoomState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLiveRoomBaseInfoRoomState) UnmarshalJSON(b []byte) error {
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

type SmartLiveRoomBaseInfoViewMode struct {
	value string
}

type SmartLiveRoomBaseInfoViewModeEnum struct {
	LANDSCAPE SmartLiveRoomBaseInfoViewMode
	VERTICAL  SmartLiveRoomBaseInfoViewMode
}

func GetSmartLiveRoomBaseInfoViewModeEnum() SmartLiveRoomBaseInfoViewModeEnum {
	return SmartLiveRoomBaseInfoViewModeEnum{
		LANDSCAPE: SmartLiveRoomBaseInfoViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: SmartLiveRoomBaseInfoViewMode{
			value: "VERTICAL",
		},
	}
}

func (c SmartLiveRoomBaseInfoViewMode) Value() string {
	return c.value
}

func (c SmartLiveRoomBaseInfoViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLiveRoomBaseInfoViewMode) UnmarshalJSON(b []byte) error {
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

type SmartLiveRoomBaseInfoLastJobStatus struct {
	value string
}

type SmartLiveRoomBaseInfoLastJobStatusEnum struct {
	WAITING    SmartLiveRoomBaseInfoLastJobStatus
	PROCESSING SmartLiveRoomBaseInfoLastJobStatus
	SUCCEED    SmartLiveRoomBaseInfoLastJobStatus
	FAILED     SmartLiveRoomBaseInfoLastJobStatus
	CANCELED   SmartLiveRoomBaseInfoLastJobStatus
	BLOCKED    SmartLiveRoomBaseInfoLastJobStatus
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
		BLOCKED: SmartLiveRoomBaseInfoLastJobStatus{
			value: "BLOCKED",
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
