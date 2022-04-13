package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSpResourceResponse struct {
	Te1080pHardCount *ResDetailDto `json:"te1080pHardCount,omitempty"`

	Te720pHardCount *ResDetailDto `json:"te720pHardCount,omitempty"`

	TeSoftCount *ResDetailDto `json:"teSoftCount,omitempty"`

	RoomCount *ResDetailDto `json:"roomCount,omitempty"`

	RecordCapability *ResDetailDto `json:"recordCapability,omitempty"`

	ConfCallCount *ResDetailDto `json:"confCallCount,omitempty"`

	LiveCount *ResDetailDto `json:"liveCount,omitempty"`

	CorpCount *ResDetailDto `json:"corpCount,omitempty"`

	ThirdPartyHardCount *ResDetailDto `json:"thirdPartyHardCount,omitempty"`

	HwVisionCount *ResDetailDto `json:"hwVisionCount,omitempty"`

	IdeaHubCount *ResDetailDto `json:"ideaHubCount,omitempty"`
	// 在创建SP的时候设置的pstn权限开关

	EnablePstn *bool `json:"enablePstn,omitempty"`
	// 在创建SP的时候设置发送短信开关

	EnableSMS *bool `json:"enableSMS,omitempty"`
	// sp管理员绑定的分组列表

	GroupList      *[]QueryCorpGroupDto `json:"groupList,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowSpResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowSpResourceResponse", string(data)}, " ")
}
