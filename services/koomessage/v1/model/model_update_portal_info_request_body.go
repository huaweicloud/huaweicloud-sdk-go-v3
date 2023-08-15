package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortalInfoRequestBody 更新主页请求体。
type UpdatePortalInfoRequestBody struct {

	// 修改原因。
	ChangeReason string `json:"change_reason"`

	// 主页背景图片资源ID。 > 分辨率大于等于1440*810，支持jpg、jpeg、bmp、png。参数值为上传智能信息服务号图片资源API返回的resource_id。
	BackgroundImg string `json:"background_img"`

	// 简介。
	Summary string `json:"summary"`

	// 热线号列表。
	Tels *[]TelModel `json:"tels,omitempty"`

	// 快应用列表。
	Fastapps *[]UpdatePortalFastappModel `json:"fastapps,omitempty"`

	// 华为服务号列表。  > 预留，暂未使用。
	HwPubs *[]string `json:"hw_pubs,omitempty"`
}

func (o UpdatePortalInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortalInfoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePortalInfoRequestBody", string(data)}, " ")
}
