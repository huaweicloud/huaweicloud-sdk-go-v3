package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSpResourceResponse struct {
	Te1080pHardCount *ResDetailDto `json:"te1080pHardCount,omitempty" xml:"te1080pHardCount"`

	Te720pHardCount *ResDetailDto `json:"te720pHardCount,omitempty" xml:"te720pHardCount"`

	TeSoftCount *ResDetailDto `json:"teSoftCount,omitempty" xml:"teSoftCount"`

	RoomCount *ResDetailDto `json:"roomCount,omitempty" xml:"roomCount"`

	RecordCapability *ResDetailDto `json:"recordCapability,omitempty" xml:"recordCapability"`

	ConfCallCount *ResDetailDto `json:"confCallCount,omitempty" xml:"confCallCount"`

	LiveCount *ResDetailDto `json:"liveCount,omitempty" xml:"liveCount"`

	CorpCount *ResDetailDto `json:"corpCount,omitempty" xml:"corpCount"`

	ThirdPartyHardCount *ResDetailDto `json:"thirdPartyHardCount,omitempty" xml:"thirdPartyHardCount"`

	HwVisionCount *ResDetailDto `json:"hwVisionCount,omitempty" xml:"hwVisionCount"`

	IdeaHubCount *ResDetailDto `json:"ideaHubCount,omitempty" xml:"ideaHubCount"`

	// 在创建SP的时候设置的pstn权限开关
	EnablePstn *bool `json:"enablePstn,omitempty" xml:"enablePstn"`

	// 在创建SP的时候设置发送短信开关
	EnableSMS *bool `json:"enableSMS,omitempty" xml:"enableSMS"`

	// sp管理员绑定的分组列表
	GroupList      *[]QueryCorpGroupDto `json:"groupList,omitempty" xml:"groupList"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowSpResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowSpResourceResponse", string(data)}, " ")
}
