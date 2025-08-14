package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DigitalAssetInfo 资产信息。
type DigitalAssetInfo struct {

	// 租户id
	ProjectId *string `json:"project_id,omitempty"`

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// ai标识ID。
	ProduceId *string `json:"produce_id,omitempty"`

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

	// 资产类型。  公共资产类型： * VOICE_MODEL：音色模型 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产类型： * HUMAN_MODEL_2D：分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板
	AssetType *DigitalAssetInfoAssetType `json:"asset_type,omitempty"`

	// 资产状态。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复 * BLOCK: 资产被冻结，资产不可用，不可查看文件。 * WAITING_DELETE：资产将被下线
	AssetState *DigitalAssetInfoAssetState `json:"asset_state,omitempty"`

	// 失败原因。 * AUTOMATIC_REVIEW_REJECT：自动审核失败 * MANUAL_REVIEW_REJECT：人工审核失败
	FailType *DigitalAssetInfoFailType `json:"fail_type,omitempty"`

	// 冻结原因编号。
	BlockReasonCode *string `json:"block_reason_code,omitempty"`

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

	// 资产自动处理任务。
	AutoOperationConfig *[]AutoOperationConfig `json:"auto_operation_config,omitempty"`
}

func (o DigitalAssetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DigitalAssetInfo struct{}"
	}

	return strings.Join([]string{"DigitalAssetInfo", string(data)}, " ")
}

type DigitalAssetInfoAssetType struct {
	value string
}

type DigitalAssetInfoAssetTypeEnum struct {
	HUMAN_MODEL           DigitalAssetInfoAssetType
	VOICE_MODEL           DigitalAssetInfoAssetType
	SCENE                 DigitalAssetInfoAssetType
	ANIMATION             DigitalAssetInfoAssetType
	VIDEO                 DigitalAssetInfoAssetType
	IMAGE                 DigitalAssetInfoAssetType
	PPT                   DigitalAssetInfoAssetType
	MATERIAL              DigitalAssetInfoAssetType
	NORMAL_MODEL          DigitalAssetInfoAssetType
	COMMON_FILE           DigitalAssetInfoAssetType
	HUMAN_MODEL_2_D       DigitalAssetInfoAssetType
	BUSINESS_CARD_TEMPLET DigitalAssetInfoAssetType
	MUSIC                 DigitalAssetInfoAssetType
	AUDIO                 DigitalAssetInfoAssetType
}

func GetDigitalAssetInfoAssetTypeEnum() DigitalAssetInfoAssetTypeEnum {
	return DigitalAssetInfoAssetTypeEnum{
		HUMAN_MODEL: DigitalAssetInfoAssetType{
			value: "HUMAN_MODEL",
		},
		VOICE_MODEL: DigitalAssetInfoAssetType{
			value: "VOICE_MODEL",
		},
		SCENE: DigitalAssetInfoAssetType{
			value: "SCENE",
		},
		ANIMATION: DigitalAssetInfoAssetType{
			value: "ANIMATION",
		},
		VIDEO: DigitalAssetInfoAssetType{
			value: "VIDEO",
		},
		IMAGE: DigitalAssetInfoAssetType{
			value: "IMAGE",
		},
		PPT: DigitalAssetInfoAssetType{
			value: "PPT",
		},
		MATERIAL: DigitalAssetInfoAssetType{
			value: "MATERIAL",
		},
		NORMAL_MODEL: DigitalAssetInfoAssetType{
			value: "NORMAL_MODEL",
		},
		COMMON_FILE: DigitalAssetInfoAssetType{
			value: "COMMON_FILE",
		},
		HUMAN_MODEL_2_D: DigitalAssetInfoAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: DigitalAssetInfoAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
		MUSIC: DigitalAssetInfoAssetType{
			value: "MUSIC",
		},
		AUDIO: DigitalAssetInfoAssetType{
			value: "AUDIO",
		},
	}
}

func (c DigitalAssetInfoAssetType) Value() string {
	return c.value
}

func (c DigitalAssetInfoAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetInfoAssetType) UnmarshalJSON(b []byte) error {
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

type DigitalAssetInfoAssetState struct {
	value string
}

type DigitalAssetInfoAssetStateEnum struct {
	CREATING       DigitalAssetInfoAssetState
	FAILED         DigitalAssetInfoAssetState
	UNACTIVED      DigitalAssetInfoAssetState
	ACTIVED        DigitalAssetInfoAssetState
	DELETING       DigitalAssetInfoAssetState
	DELETED        DigitalAssetInfoAssetState
	BLOCK          DigitalAssetInfoAssetState
	WAITING_DELETE DigitalAssetInfoAssetState
}

func GetDigitalAssetInfoAssetStateEnum() DigitalAssetInfoAssetStateEnum {
	return DigitalAssetInfoAssetStateEnum{
		CREATING: DigitalAssetInfoAssetState{
			value: "CREATING",
		},
		FAILED: DigitalAssetInfoAssetState{
			value: "FAILED",
		},
		UNACTIVED: DigitalAssetInfoAssetState{
			value: "UNACTIVED",
		},
		ACTIVED: DigitalAssetInfoAssetState{
			value: "ACTIVED",
		},
		DELETING: DigitalAssetInfoAssetState{
			value: "DELETING",
		},
		DELETED: DigitalAssetInfoAssetState{
			value: "DELETED",
		},
		BLOCK: DigitalAssetInfoAssetState{
			value: "BLOCK",
		},
		WAITING_DELETE: DigitalAssetInfoAssetState{
			value: "WAITING_DELETE",
		},
	}
}

func (c DigitalAssetInfoAssetState) Value() string {
	return c.value
}

func (c DigitalAssetInfoAssetState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetInfoAssetState) UnmarshalJSON(b []byte) error {
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

type DigitalAssetInfoFailType struct {
	value string
}

type DigitalAssetInfoFailTypeEnum struct {
	AUTOMATIC_REVIEW_REJECT DigitalAssetInfoFailType
	MANUAL_REVIEW_REJECT    DigitalAssetInfoFailType
}

func GetDigitalAssetInfoFailTypeEnum() DigitalAssetInfoFailTypeEnum {
	return DigitalAssetInfoFailTypeEnum{
		AUTOMATIC_REVIEW_REJECT: DigitalAssetInfoFailType{
			value: "AUTOMATIC_REVIEW_REJECT",
		},
		MANUAL_REVIEW_REJECT: DigitalAssetInfoFailType{
			value: "MANUAL_REVIEW_REJECT",
		},
	}
}

func (c DigitalAssetInfoFailType) Value() string {
	return c.value
}

func (c DigitalAssetInfoFailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetInfoFailType) UnmarshalJSON(b []byte) error {
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
