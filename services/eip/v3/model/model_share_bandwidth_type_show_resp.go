package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShareBandwidthTypeShowResp 带宽支持类型对象
type ShareBandwidthTypeShowResp struct {

	// 支持带宽类型的id
	Id *string `json:"id,omitempty"`

	// 带宽类型
	BandwidthType *string `json:"bandwidth_type,omitempty"`

	// 中心站点or边缘站点，默认展示
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 带宽类型的英文表述
	NameEn *string `json:"name_en,omitempty"`

	// 带宽类型的中文表述
	NameZh *string `json:"name_zh,omitempty"`

	// 带宽类型描述信息
	Description *string `json:"description,omitempty"`
}

func (o ShareBandwidthTypeShowResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareBandwidthTypeShowResp struct{}"
	}

	return strings.Join([]string{"ShareBandwidthTypeShowResp", string(data)}, " ")
}
