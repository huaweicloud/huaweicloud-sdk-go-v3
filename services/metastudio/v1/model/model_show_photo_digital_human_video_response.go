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

	// 任务的状态。 * WAITING：等待 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败 * CANCELED：取消
	State ShowPhotoDigitalHumanVideoResponseState `json:"state"`

	// 数字人视频制作开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人视频制作结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 数字人视频内容时长。
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