package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDigitalAssetResponse Response Object
type UpdateDigitalAssetResponse struct {

	// 租户id
	ProjectId *string `json:"project_id,omitempty"`

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产描述。
	AssetDescription *string `json:"asset_description,omitempty"`

	// 第三方用户ID。 > * 即创建资产是通过X-App-UserId头域传入的值。
	AppUserId *string `json:"app_user_id,omitempty"`

	// 资产创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 资产更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 资产类型。  公共资产类型： * VOICE_MODEL：音色模型 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产类型： * HUMAN_MODEL_2D：分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产类型： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型
	AssetType *UpdateDigitalAssetResponseAssetType `json:"asset_type,omitempty"`

	// 资产状态。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复 * BLOCK: 资产被冻结，资产不可用，不可查看文件。 * WAITING_DELETE：资产将被下线
	AssetState *UpdateDigitalAssetResponseAssetState `json:"asset_state,omitempty"`

	// 失败原因。 * AUTOMATIC_REVIEW_REJECT：自动审核失败 * MANUAL_REVIEW_REJECT：人工审核失败
	FailType *UpdateDigitalAssetResponseFailType `json:"fail_type,omitempty"`

	// 冻结/解冻/失败 原因。
	Reason *string `json:"reason,omitempty"`

	// 标签列表。 > 分身形象系统资产的tag定义如下： > - 行业：NEWS,BUSINESS,E-COMMERCE,MARKETING,KNOWLEDGE,EDUCATION,SPORTS > - 性别：MALE,FEMALE > - 姿势：FULL-BODY,HALF-BODY,STANDING,SITTING,WALKING > - 区域：ASIAN,WESTERN,MIDDLE-EASTERNER,AFRICAN,LATINO
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`

	// 资产下的文件。
	Files *[]AssetFileInfo `json:"files,omitempty"`

	// 展示顺序
	AssetOrder *int32 `json:"asset_order,omitempty"`

	// 支持的业务类型。： * VIDEO_2D：分身数字人视频制作 * LIVE_2D：分身数字人直播 * CHAT_2D：分身数字人智能交互
	SupportedService *[]SupportedServiceEnum `json:"supported_service,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDigitalAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDigitalAssetResponse struct{}"
	}

	return strings.Join([]string{"UpdateDigitalAssetResponse", string(data)}, " ")
}

type UpdateDigitalAssetResponseAssetType struct {
	value string
}

type UpdateDigitalAssetResponseAssetTypeEnum struct {
	HUMAN_MODEL           UpdateDigitalAssetResponseAssetType
	VOICE_MODEL           UpdateDigitalAssetResponseAssetType
	SCENE                 UpdateDigitalAssetResponseAssetType
	ANIMATION             UpdateDigitalAssetResponseAssetType
	VIDEO                 UpdateDigitalAssetResponseAssetType
	IMAGE                 UpdateDigitalAssetResponseAssetType
	PPT                   UpdateDigitalAssetResponseAssetType
	MATERIAL              UpdateDigitalAssetResponseAssetType
	NORMAL_MODEL          UpdateDigitalAssetResponseAssetType
	COMMON_FILE           UpdateDigitalAssetResponseAssetType
	HUMAN_MODEL_2_D       UpdateDigitalAssetResponseAssetType
	BUSINESS_CARD_TEMPLET UpdateDigitalAssetResponseAssetType
	MUSIC                 UpdateDigitalAssetResponseAssetType
	AUDIO                 UpdateDigitalAssetResponseAssetType
}

func GetUpdateDigitalAssetResponseAssetTypeEnum() UpdateDigitalAssetResponseAssetTypeEnum {
	return UpdateDigitalAssetResponseAssetTypeEnum{
		HUMAN_MODEL: UpdateDigitalAssetResponseAssetType{
			value: "HUMAN_MODEL",
		},
		VOICE_MODEL: UpdateDigitalAssetResponseAssetType{
			value: "VOICE_MODEL",
		},
		SCENE: UpdateDigitalAssetResponseAssetType{
			value: "SCENE",
		},
		ANIMATION: UpdateDigitalAssetResponseAssetType{
			value: "ANIMATION",
		},
		VIDEO: UpdateDigitalAssetResponseAssetType{
			value: "VIDEO",
		},
		IMAGE: UpdateDigitalAssetResponseAssetType{
			value: "IMAGE",
		},
		PPT: UpdateDigitalAssetResponseAssetType{
			value: "PPT",
		},
		MATERIAL: UpdateDigitalAssetResponseAssetType{
			value: "MATERIAL",
		},
		NORMAL_MODEL: UpdateDigitalAssetResponseAssetType{
			value: "NORMAL_MODEL",
		},
		COMMON_FILE: UpdateDigitalAssetResponseAssetType{
			value: "COMMON_FILE",
		},
		HUMAN_MODEL_2_D: UpdateDigitalAssetResponseAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: UpdateDigitalAssetResponseAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
		MUSIC: UpdateDigitalAssetResponseAssetType{
			value: "MUSIC",
		},
		AUDIO: UpdateDigitalAssetResponseAssetType{
			value: "AUDIO",
		},
	}
}

func (c UpdateDigitalAssetResponseAssetType) Value() string {
	return c.value
}

func (c UpdateDigitalAssetResponseAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDigitalAssetResponseAssetType) UnmarshalJSON(b []byte) error {
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

type UpdateDigitalAssetResponseAssetState struct {
	value string
}

type UpdateDigitalAssetResponseAssetStateEnum struct {
	CREATING       UpdateDigitalAssetResponseAssetState
	FAILED         UpdateDigitalAssetResponseAssetState
	UNACTIVED      UpdateDigitalAssetResponseAssetState
	ACTIVED        UpdateDigitalAssetResponseAssetState
	DELETING       UpdateDigitalAssetResponseAssetState
	DELETED        UpdateDigitalAssetResponseAssetState
	BLOCK          UpdateDigitalAssetResponseAssetState
	WAITING_DELETE UpdateDigitalAssetResponseAssetState
}

func GetUpdateDigitalAssetResponseAssetStateEnum() UpdateDigitalAssetResponseAssetStateEnum {
	return UpdateDigitalAssetResponseAssetStateEnum{
		CREATING: UpdateDigitalAssetResponseAssetState{
			value: "CREATING",
		},
		FAILED: UpdateDigitalAssetResponseAssetState{
			value: "FAILED",
		},
		UNACTIVED: UpdateDigitalAssetResponseAssetState{
			value: "UNACTIVED",
		},
		ACTIVED: UpdateDigitalAssetResponseAssetState{
			value: "ACTIVED",
		},
		DELETING: UpdateDigitalAssetResponseAssetState{
			value: "DELETING",
		},
		DELETED: UpdateDigitalAssetResponseAssetState{
			value: "DELETED",
		},
		BLOCK: UpdateDigitalAssetResponseAssetState{
			value: "BLOCK",
		},
		WAITING_DELETE: UpdateDigitalAssetResponseAssetState{
			value: "WAITING_DELETE",
		},
	}
}

func (c UpdateDigitalAssetResponseAssetState) Value() string {
	return c.value
}

func (c UpdateDigitalAssetResponseAssetState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDigitalAssetResponseAssetState) UnmarshalJSON(b []byte) error {
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

type UpdateDigitalAssetResponseFailType struct {
	value string
}

type UpdateDigitalAssetResponseFailTypeEnum struct {
	AUTOMATIC_REVIEW_REJECT UpdateDigitalAssetResponseFailType
	MANUAL_REVIEW_REJECT    UpdateDigitalAssetResponseFailType
}

func GetUpdateDigitalAssetResponseFailTypeEnum() UpdateDigitalAssetResponseFailTypeEnum {
	return UpdateDigitalAssetResponseFailTypeEnum{
		AUTOMATIC_REVIEW_REJECT: UpdateDigitalAssetResponseFailType{
			value: "AUTOMATIC_REVIEW_REJECT",
		},
		MANUAL_REVIEW_REJECT: UpdateDigitalAssetResponseFailType{
			value: "MANUAL_REVIEW_REJECT",
		},
	}
}

func (c UpdateDigitalAssetResponseFailType) Value() string {
	return c.value
}

func (c UpdateDigitalAssetResponseFailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDigitalAssetResponseFailType) UnmarshalJSON(b []byte) error {
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
