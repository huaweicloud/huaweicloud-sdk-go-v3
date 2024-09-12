package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Show2DDigitalHumanVideoResponse Response Object
type Show2DDigitalHumanVideoResponse struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING：等待 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败 * CANCELED：取消 * BLOCK: 冻结
	State Show2DDigitalHumanVideoResponseState `json:"state"`

	// 任务类型。 * 2D_DIGITAL_HUMAN_VIDEO: 分身数字人视频制作任务 * PHOTO_DIGITAL_HUMAN_VIDEO: 照片数字人视频制作任务
	JobType *Show2DDigitalHumanVideoResponseJobType `json:"job_type,omitempty"`

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

	// 视频生成类型。该参数取值是MODEL时，model_asset_id必填；取值是PICTURE时，human_image必填。 * MODEL：通过分身数字人模型生成视频 * PICTURE： 通过单张照片生成视频 > * 该参数已废弃，照片数字人视频制作使用“创建照片分身数字人视频制作任务”接口。
	VideoMakingType *Show2DDigitalHumanVideoResponseVideoMakingType `json:"video_making_type,omitempty"`

	// 人物照片，需要Base64编码。 > * 该参数已废弃，照片数字人视频制作使用“创建照片分身数字人视频制作任务”接口。
	HumanImage *string `json:"human_image,omitempty"`

	// 分身数字人模型资产ID，可以从资产库中查询。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptItem `json:"shoot_scripts,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Show2DDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Show2DDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"Show2DDigitalHumanVideoResponse", string(data)}, " ")
}

type Show2DDigitalHumanVideoResponseState struct {
	value string
}

type Show2DDigitalHumanVideoResponseStateEnum struct {
	WAITING    Show2DDigitalHumanVideoResponseState
	PROCESSING Show2DDigitalHumanVideoResponseState
	SUCCEED    Show2DDigitalHumanVideoResponseState
	FAILED     Show2DDigitalHumanVideoResponseState
	CANCELED   Show2DDigitalHumanVideoResponseState
	BLOCK      Show2DDigitalHumanVideoResponseState
}

func GetShow2DDigitalHumanVideoResponseStateEnum() Show2DDigitalHumanVideoResponseStateEnum {
	return Show2DDigitalHumanVideoResponseStateEnum{
		WAITING: Show2DDigitalHumanVideoResponseState{
			value: "WAITING",
		},
		PROCESSING: Show2DDigitalHumanVideoResponseState{
			value: "PROCESSING",
		},
		SUCCEED: Show2DDigitalHumanVideoResponseState{
			value: "SUCCEED",
		},
		FAILED: Show2DDigitalHumanVideoResponseState{
			value: "FAILED",
		},
		CANCELED: Show2DDigitalHumanVideoResponseState{
			value: "CANCELED",
		},
		BLOCK: Show2DDigitalHumanVideoResponseState{
			value: "BLOCK",
		},
	}
}

func (c Show2DDigitalHumanVideoResponseState) Value() string {
	return c.value
}

func (c Show2DDigitalHumanVideoResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2DDigitalHumanVideoResponseState) UnmarshalJSON(b []byte) error {
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

type Show2DDigitalHumanVideoResponseJobType struct {
	value string
}

type Show2DDigitalHumanVideoResponseJobTypeEnum struct {
	E_2_D_DIGITAL_HUMAN_VIDEO Show2DDigitalHumanVideoResponseJobType
	DIGITAL_HUMAN_PHOTO_VIDEO Show2DDigitalHumanVideoResponseJobType
}

func GetShow2DDigitalHumanVideoResponseJobTypeEnum() Show2DDigitalHumanVideoResponseJobTypeEnum {
	return Show2DDigitalHumanVideoResponseJobTypeEnum{
		E_2_D_DIGITAL_HUMAN_VIDEO: Show2DDigitalHumanVideoResponseJobType{
			value: "2D_DIGITAL_HUMAN_VIDEO",
		},
		DIGITAL_HUMAN_PHOTO_VIDEO: Show2DDigitalHumanVideoResponseJobType{
			value: "DIGITAL_HUMAN_PHOTO_VIDEO",
		},
	}
}

func (c Show2DDigitalHumanVideoResponseJobType) Value() string {
	return c.value
}

func (c Show2DDigitalHumanVideoResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2DDigitalHumanVideoResponseJobType) UnmarshalJSON(b []byte) error {
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

type Show2DDigitalHumanVideoResponseVideoMakingType struct {
	value string
}

type Show2DDigitalHumanVideoResponseVideoMakingTypeEnum struct {
	MODEL   Show2DDigitalHumanVideoResponseVideoMakingType
	PICTURE Show2DDigitalHumanVideoResponseVideoMakingType
}

func GetShow2DDigitalHumanVideoResponseVideoMakingTypeEnum() Show2DDigitalHumanVideoResponseVideoMakingTypeEnum {
	return Show2DDigitalHumanVideoResponseVideoMakingTypeEnum{
		MODEL: Show2DDigitalHumanVideoResponseVideoMakingType{
			value: "MODEL",
		},
		PICTURE: Show2DDigitalHumanVideoResponseVideoMakingType{
			value: "PICTURE",
		},
	}
}

func (c Show2DDigitalHumanVideoResponseVideoMakingType) Value() string {
	return c.value
}

func (c Show2DDigitalHumanVideoResponseVideoMakingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2DDigitalHumanVideoResponseVideoMakingType) UnmarshalJSON(b []byte) error {
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
