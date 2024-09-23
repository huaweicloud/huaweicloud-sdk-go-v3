package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAssetsRequest Request Object
type ListAssetsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 按名称模糊查询。
	Name *string `json:"name,omitempty"`

	// 按标签模糊查询。
	Tag *string `json:"tag,omitempty"`

	// 标签查询组合方式 INTERSECTION：交集 UNION_SET：并集
	TagCombinationType *ListAssetsRequestTagCombinationType `json:"tag_combination_type,omitempty"`

	// 最近直播任务起始时间。格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	// 资产类型。多个类型使用英文逗号分割。 * HUMAN_MODEL：数字人模型 * VOICE_MODEL：音色模型（仅系统管理员可上传） * SCENE：场景模型 * ANIMATION：动作动画 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MATERIAL：风格化素材 * HUMAN_MODEL_2D: 2D数字人网络模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板 * MUSIC: 音乐 * AUDIO: 音频
	AssetType *string `json:"asset_type,omitempty"`

	// 排序字段，支持的排序方式有： - 按创建时间排序：create_time - 按更新时间排序：update_time - 按资产排序：asset_order
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式。 * asc：升序 * desc：降序  默认asc升序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 资产来源。 * SYSTEM：系统资产 * CUSTOMIZATION：租户资产 * ALL：所有资产  默认查询租户资产。
	AssetSource *ListAssetsRequestAssetSource `json:"asset_source,omitempty"`

	// 资产状态。多个资产状态使用英文逗号分割。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复 * BLOCK：资产被冻结，资产不可用，不可查看文件。 * WAITING_DELETE：资产将被下线 默认查询所有状态的资产。
	AssetState *string `json:"asset_state,omitempty"`

	// 基于风格化ID查询关联资产。 * system_male_001：男性风格01 * system_female_001：女性风格01 * system_male_002：男性风格02  * system_female_002：女性风格02
	StyleId *string `json:"style_id,omitempty"`

	// 使用精确查询的字段
	AccurateQueryField *[]string `json:"accurate_query_field,omitempty"`

	// 可用引擎。 * UE：UE引擎 * MetaEngine：MetaEngine引擎 > 该字段当前只对MetaEngine白名单用户生效
	RenderEngine *string `json:"render_engine,omitempty"`

	// 资产id
	AssetId *[]string `json:"asset_id,omitempty"`

	// 性别。多选使用英文逗号分隔。
	Sex *string `json:"sex,omitempty"`

	// 语言。多选使用英文逗号分隔。
	Language *string `json:"language,omitempty"`

	// 系统属性。  key和value间用\":\"分隔，多个key之间用\",\"分隔。  如system_property=BACKGROUND_IMG:Yes,RENDER_ENGINE:MetaEngine。  不同Key对应Value取值如下：  公共资产属性： * BACKGROUND_IMG：视频制作的2D背景图片，可取值Yes * CREATED_BY_PLATFORM：是否平台生成，可取值Yes  分身数字人资产属性： * MATERIAL_IMG：素材图片，用作前景。可取值Yes * MATERIAL_VIDEO：素材视频，用作前景。可取值Yes * TO_BE_TRANSLATED_VIDEO: 视频翻译的源视频。可取值Yes  3D数字人资产属性： * STYLE_ID：风格Id * RENDER_ENGINE：引擎类型，可取值UE或MetaEngine * BACKGROUND_SCENE：视频制作的2D背景场景，可取值Horizontal（横屏）或者Vertical（竖屏）
	SystemProperty *string `json:"system_property,omitempty"`

	// 动作是否可编辑。仅在分身数字人模型查询时有效。
	ActionEditable *bool `json:"action_editable,omitempty"`

	// 分身数字人是否带原子动作库。 > * 带原子动作库的分身数字人可做动作编排。
	IsWithActionLibrary *bool `json:"is_with_action_library,omitempty"`

	// 分身数字人是否支持走动。仅在分身数字人模型查询时有效。
	IsMovable *bool `json:"is_movable,omitempty"`

	// 取值：HUAWEI_METASTUDIO、MOBVOI。 HUAWEI_METASTUDIO：MetaStudio自研音色 MOBVOI：出门问问音色
	VoiceProvider *string `json:"voice_provider,omitempty"`

	// 角色。 SHARER：共享方，SHAREE：被共享方
	Role *ListAssetsRequestRole `json:"role,omitempty"`

	// 音色是否支持实时合成。仅在音色查询时有效。 > * 支持实时合成的音色，可以用于直播和智能交互场景。否则只能用于视频制作。
	IsRealtimeVoice *bool `json:"is_realtime_voice,omitempty"`

	// 模型版本
	HumanModel2dVersion *string `json:"human_model_2d_version,omitempty"`

	// 资产已执行的任务名称
	IncludeDeviceName *string `json:"include_device_name,omitempty"`

	// 资产已执行的任务名称
	ExcludeDeviceName *string `json:"exclude_device_name,omitempty"`

	// 资产支持的业务类型。默认查询所有资产。 * VIDEO_2D：分身数字人视频制作 * LIVE_2D：分身数字人直播 * CHAT_2D：分身数字人智能交互
	SupportedService *ListAssetsRequestSupportedService `json:"supported_service,omitempty"`
}

func (o ListAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetsRequest struct{}"
	}

	return strings.Join([]string{"ListAssetsRequest", string(data)}, " ")
}

type ListAssetsRequestTagCombinationType struct {
	value string
}

type ListAssetsRequestTagCombinationTypeEnum struct {
	INTERSECTION ListAssetsRequestTagCombinationType
	UNION_SET    ListAssetsRequestTagCombinationType
}

func GetListAssetsRequestTagCombinationTypeEnum() ListAssetsRequestTagCombinationTypeEnum {
	return ListAssetsRequestTagCombinationTypeEnum{
		INTERSECTION: ListAssetsRequestTagCombinationType{
			value: "INTERSECTION",
		},
		UNION_SET: ListAssetsRequestTagCombinationType{
			value: "UNION_SET",
		},
	}
}

func (c ListAssetsRequestTagCombinationType) Value() string {
	return c.value
}

func (c ListAssetsRequestTagCombinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssetsRequestTagCombinationType) UnmarshalJSON(b []byte) error {
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

type ListAssetsRequestAssetSource struct {
	value string
}

type ListAssetsRequestAssetSourceEnum struct {
	SYSTEM        ListAssetsRequestAssetSource
	CUSTOMIZATION ListAssetsRequestAssetSource
	ALL           ListAssetsRequestAssetSource
}

func GetListAssetsRequestAssetSourceEnum() ListAssetsRequestAssetSourceEnum {
	return ListAssetsRequestAssetSourceEnum{
		SYSTEM: ListAssetsRequestAssetSource{
			value: "SYSTEM",
		},
		CUSTOMIZATION: ListAssetsRequestAssetSource{
			value: "CUSTOMIZATION",
		},
		ALL: ListAssetsRequestAssetSource{
			value: "ALL",
		},
	}
}

func (c ListAssetsRequestAssetSource) Value() string {
	return c.value
}

func (c ListAssetsRequestAssetSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssetsRequestAssetSource) UnmarshalJSON(b []byte) error {
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

type ListAssetsRequestRole struct {
	value string
}

type ListAssetsRequestRoleEnum struct {
	SHARER ListAssetsRequestRole
	SHAREE ListAssetsRequestRole
}

func GetListAssetsRequestRoleEnum() ListAssetsRequestRoleEnum {
	return ListAssetsRequestRoleEnum{
		SHARER: ListAssetsRequestRole{
			value: "SHARER",
		},
		SHAREE: ListAssetsRequestRole{
			value: "SHAREE",
		},
	}
}

func (c ListAssetsRequestRole) Value() string {
	return c.value
}

func (c ListAssetsRequestRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssetsRequestRole) UnmarshalJSON(b []byte) error {
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

type ListAssetsRequestSupportedService struct {
	value string
}

type ListAssetsRequestSupportedServiceEnum struct {
	VIDEO_2_D ListAssetsRequestSupportedService
	LIVE_2_D  ListAssetsRequestSupportedService
	CHAT_2_D  ListAssetsRequestSupportedService
}

func GetListAssetsRequestSupportedServiceEnum() ListAssetsRequestSupportedServiceEnum {
	return ListAssetsRequestSupportedServiceEnum{
		VIDEO_2_D: ListAssetsRequestSupportedService{
			value: "VIDEO_2D",
		},
		LIVE_2_D: ListAssetsRequestSupportedService{
			value: "LIVE_2D",
		},
		CHAT_2_D: ListAssetsRequestSupportedService{
			value: "CHAT_2D",
		},
	}
}

func (c ListAssetsRequestSupportedService) Value() string {
	return c.value
}

func (c ListAssetsRequestSupportedService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssetsRequestSupportedService) UnmarshalJSON(b []byte) error {
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
