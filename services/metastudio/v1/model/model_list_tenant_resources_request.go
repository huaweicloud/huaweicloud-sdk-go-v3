package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantResourcesRequest Request Object
type ListTenantResourcesRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 资源类型。可填多个，用\",\"分隔。详见[资源类型](metastudio_02_0042.xml)。
	Resource *string `json:"resource,omitempty"`

	// 业务类型。可填多个，用\",\"分隔。 * VOICE_CLONE：声音制作 * SYNTHETICS_SOUND：声音合成 * ASSET_MANAGER：资产管理 * MODELING_2D：形象制作 * LIVE_2D：分身数字人视频直播 * VIDEO_2D：分身数字人视频制作 * CHAT_2D：分身数字人智能交互 * BUSINESS_CARD_2D：分身数字人名片 * PICTURE_2D：照片数字人视频 * MODELING_3D：3D照片建模 * VDS_3D：3D视觉驱动 * TTSA_3D：3D语音驱动 * FLEXUS_2D：FLEXUS版本资源
	Business *string `json:"business,omitempty"`

	// 资源来源,可填多个 例如:PURCHASED,ADMIN_ALLOCATED,将返回商用资源与管理员分配资源。 * PURCHASED: 用户购买的资源 * SP_ALLOCATED: SP分配的资源 * ADMIN_ALLOCATED: 系统管理员分配的资源
	ResourceSource string `json:"resource_source"`

	// 资源id。
	ResourceId *string `json:"resource_id,omitempty"`

	// 订单id。
	OrderId *string `json:"order_id,omitempty"`

	// 计费类型。  * PERIODIC: 包周期  * ONE_TIME：一次性  * ON_DEMAND：按需 计费模式。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源过期时间段 开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireStartTime *string `json:"resource_expire_start_time,omitempty"`

	// 资源过期时间段 结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireEndTime *string `json:"resource_expire_end_time,omitempty"`

	// 子资源类型。当前只有flexus套餐包存在该字段 * voice_clone_flexus: 语音克隆Flexus版 * modeling_count_2d_model_flexus: 分身数字人形象制作Flexus版 * video_time_flexus_2d_model: 分身数字人Flexus版本视频制作
	SubResource *string `json:"sub_resource,omitempty"`
}

func (o ListTenantResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantResourcesRequest", string(data)}, " ")
}
