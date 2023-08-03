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

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 按名称模糊查询。
	Name *string `json:"name,omitempty"`

	// 按标签模糊查询。
	Tag *string `json:"tag,omitempty"`

	// 起始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	// 资产类型。多个类型使用英文逗号分割。 * HUMAN_MODEL：数字人模型 * VOICE_MODEL：音色模型（仅系统管理员可上传） * SCENE：场景模型 * ANIMATION：动作动画 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MATERIAL：风格化素材 * HUMAN_MODEL_2D: 2D数字人网络模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板
	AssetType *string `json:"asset_type,omitempty"`

	// 排序字段，目前只支持create_time。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式。 * asc：升序 * desc：降序  默认asc升序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 资产来源。 * SYSTEM：系统资产 * CUSTOMIZATION：租户资产 * ALL：所有资产  默认查询租户资产。
	AssetSource *ListAssetsRequestAssetSource `json:"asset_source,omitempty"`

	// 资产管理分类。 * UPLOAD：我的上传 * UPLOADED：已上传 * UPLOADING：UPLOADING * UPLOAD_FAILED：上传失败 * DOWNLOAD：我的下载 * COLLECTION：我的收藏 * DRAFT：草稿箱 * RECYCLE：回收站
	AssetManageType *ListAssetsRequestAssetManageType `json:"asset_manage_type,omitempty"`

	// 资产状态。多个资产状态使用英文逗号分割。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复  默认查询所有状态的资产。
	AssetState *string `json:"asset_state,omitempty"`

	// 基于风格化ID查询关联资产。 * system_male_001：男性风格01 * system_female_001：女性风格01 * system_male_002：男性风格02  * system_female_002：女性风格02
	StyleId *string `json:"style_id,omitempty"`

	// 可用引擎。 * UE：UE引擎 * MetaEngine：MetaEngine引擎 > 该字段当前只对MetaEngine白名单用户生效
	RenderEngine *string `json:"render_engine,omitempty"`

	// 性别。多选使用英文逗号分隔。
	Sex *string `json:"sex,omitempty"`

	// 语言。多选使用英文逗号分隔。
	Language *string `json:"language,omitempty"`

	// 系统属性。  key和value间用\":\"分隔，多个key之间用\",\"分隔。  如system_property=BACKGROUND_IMG:Yes,RENDER_ENGINE:MetaEngine。  不同Key对应Value取值如下： * STYLE_ID：风格Id * RENDER_ENGINE：引擎类型，可取值UE或MetaEngine * BACKGROUND_IMG：视频制作的2D背景图片，可取值Yes * BACKGROUND_SCENE：视频制作的2D背景场景，可取值Horizontal（横屏）或者Vertical（竖屏）
	SystemProperty *string `json:"system_property,omitempty"`
}

func (o ListAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetsRequest struct{}"
	}

	return strings.Join([]string{"ListAssetsRequest", string(data)}, " ")
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

type ListAssetsRequestAssetManageType struct {
	value string
}

type ListAssetsRequestAssetManageTypeEnum struct {
	UPLOAD        ListAssetsRequestAssetManageType
	UPLOADED      ListAssetsRequestAssetManageType
	UPLOADING     ListAssetsRequestAssetManageType
	UPLOAD_FAILED ListAssetsRequestAssetManageType
	DOWNLOAD      ListAssetsRequestAssetManageType
	COLLECTIO     ListAssetsRequestAssetManageType
	DRAFT         ListAssetsRequestAssetManageType
	RECYCLE       ListAssetsRequestAssetManageType
}

func GetListAssetsRequestAssetManageTypeEnum() ListAssetsRequestAssetManageTypeEnum {
	return ListAssetsRequestAssetManageTypeEnum{
		UPLOAD: ListAssetsRequestAssetManageType{
			value: "UPLOAD",
		},
		UPLOADED: ListAssetsRequestAssetManageType{
			value: "UPLOADED",
		},
		UPLOADING: ListAssetsRequestAssetManageType{
			value: "UPLOADING",
		},
		UPLOAD_FAILED: ListAssetsRequestAssetManageType{
			value: "UPLOAD_FAILED",
		},
		DOWNLOAD: ListAssetsRequestAssetManageType{
			value: "DOWNLOAD",
		},
		COLLECTIO: ListAssetsRequestAssetManageType{
			value: "COLLECTIO",
		},
		DRAFT: ListAssetsRequestAssetManageType{
			value: "DRAFT",
		},
		RECYCLE: ListAssetsRequestAssetManageType{
			value: "RECYCLE",
		},
	}
}

func (c ListAssetsRequestAssetManageType) Value() string {
	return c.value
}

func (c ListAssetsRequestAssetManageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAssetsRequestAssetManageType) UnmarshalJSON(b []byte) error {
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
