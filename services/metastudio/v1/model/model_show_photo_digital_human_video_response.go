package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPhotoDigitalHumanVideoResponse Response Object
type ShowPhotoDigitalHumanVideoResponse struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING：等待 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败 * CANCELED：取消 * BLOCK: 冻结
	State ShowPhotoDigitalHumanVideoResponseState `json:"state"`

	// 任务类型。 * 2D_DIGITAL_HUMAN_VIDEO: 分身数字人视频制作任务 * PHOTO_DIGITAL_HUMAN_VIDEO: 照片数字人视频制作任务
	JobType *ShowPhotoDigitalHumanVideoResponseJobType `json:"job_type,omitempty"`

	// 数字人视频制作开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人视频制作结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 数字人视频内容时长。
	Duration *float32 `json:"duration,omitempty"`

	OutputAssetConfig *OutputAssetInfo `json:"output_asset_config,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`

	// 剧本ID。
	ScriptId *string `json:"script_id,omitempty"`

	// 人物照片，需要Base64编码。
	HumanImage *string `json:"human_image,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *PhotoVideoConfig `json:"video_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptItem `json:"shoot_scripts,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPhotoDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPhotoDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"ShowPhotoDigitalHumanVideoResponse", string(data)}, " ")
}

type ShowPhotoDigitalHumanVideoResponseState struct {
	value string
}

type ShowPhotoDigitalHumanVideoResponseStateEnum struct {
	WAITING    ShowPhotoDigitalHumanVideoResponseState
	PROCESSING ShowPhotoDigitalHumanVideoResponseState
	SUCCEED    ShowPhotoDigitalHumanVideoResponseState
	FAILED     ShowPhotoDigitalHumanVideoResponseState
	CANCELED   ShowPhotoDigitalHumanVideoResponseState
	BLOCK      ShowPhotoDigitalHumanVideoResponseState
}

func GetShowPhotoDigitalHumanVideoResponseStateEnum() ShowPhotoDigitalHumanVideoResponseStateEnum {
	return ShowPhotoDigitalHumanVideoResponseStateEnum{
		WAITING: ShowPhotoDigitalHumanVideoResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowPhotoDigitalHumanVideoResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ShowPhotoDigitalHumanVideoResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowPhotoDigitalHumanVideoResponseState{
			value: "FAILED",
		},
		CANCELED: ShowPhotoDigitalHumanVideoResponseState{
			value: "CANCELED",
		},
		BLOCK: ShowPhotoDigitalHumanVideoResponseState{
			value: "BLOCK",
		},
	}
}

func (c ShowPhotoDigitalHumanVideoResponseState) Value() string {
	return c.value
}

func (c ShowPhotoDigitalHumanVideoResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPhotoDigitalHumanVideoResponseState) UnmarshalJSON(b []byte) error {
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

type ShowPhotoDigitalHumanVideoResponseJobType struct {
	value string
}

type ShowPhotoDigitalHumanVideoResponseJobTypeEnum struct {
	E_2_D_DIGITAL_HUMAN_VIDEO ShowPhotoDigitalHumanVideoResponseJobType
	DIGITAL_HUMAN_PHOTO_VIDEO ShowPhotoDigitalHumanVideoResponseJobType
}

func GetShowPhotoDigitalHumanVideoResponseJobTypeEnum() ShowPhotoDigitalHumanVideoResponseJobTypeEnum {
	return ShowPhotoDigitalHumanVideoResponseJobTypeEnum{
		E_2_D_DIGITAL_HUMAN_VIDEO: ShowPhotoDigitalHumanVideoResponseJobType{
			value: "2D_DIGITAL_HUMAN_VIDEO",
		},
		DIGITAL_HUMAN_PHOTO_VIDEO: ShowPhotoDigitalHumanVideoResponseJobType{
			value: "DIGITAL_HUMAN_PHOTO_VIDEO",
		},
	}
}

func (c ShowPhotoDigitalHumanVideoResponseJobType) Value() string {
	return c.value
}

func (c ShowPhotoDigitalHumanVideoResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPhotoDigitalHumanVideoResponseJobType) UnmarshalJSON(b []byte) error {
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
