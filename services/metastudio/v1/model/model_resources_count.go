package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourcesCount 资源总量信息
type ResourcesCount struct {

	// 业务类型。可填多个，用\",\"分隔。 * VOICE_CLONE：声音制作 * SYNTHETICS_SOUND：声音合成 * ASSET_MANAGER：资产管理 * MODELING_2D：形象制作 * LIVE_2D：分身数字人视频直播 * VIDEO_2D：分身数字人视频制作 * CHAT_2D：分身数字人智能交互 * BUSINESS_CARD_2D：分身数字人名片 * PICTURE_2D：照片数字人视频 * MODELING_3D：3D照片建模 * VDS_3D：3D视觉驱动 * TTSA_3D：3D语音驱动 * FLEXUS_2D：FLEXUS版本资源
	BusinessType *string `json:"business_type,omitempty"`

	// 资源总数。
	Count float32 `json:"count,omitempty"`

	// 计费类型。 * PERIODIC: 包周期 * ONE_TIME：一次性 * ON_DEMAND：按需
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源来源。 * PURCHASED: 购买 * SP_ALLOCATED：SP分配 * ADMIN_ALLOCATED：系统管理员分配
	ResourceSource *string `json:"resource_source,omitempty"`
}

func (o ResourcesCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesCount struct{}"
	}

	return strings.Join([]string{"ResourcesCount", string(data)}, " ")
}
