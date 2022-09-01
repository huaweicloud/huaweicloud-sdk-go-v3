package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCorpResourceResponse struct {
	Te1080pHardCount *ResDetailDto `json:"te1080pHardCount,omitempty" xml:"te1080pHardCount"`

	Te720pHardCount *ResDetailDto `json:"te720pHardCount,omitempty" xml:"te720pHardCount"`

	TeSoftCount *ResDetailDto `json:"teSoftCount,omitempty" xml:"teSoftCount"`

	RoomCount *ResDetailDto `json:"roomCount,omitempty" xml:"roomCount"`

	RecordCapability *ResDetailDto `json:"recordCapability,omitempty" xml:"recordCapability"`

	ConfCallCount *ResDetailDto `json:"confCallCount,omitempty" xml:"confCallCount"`

	LiveCount *ResDetailDto `json:"liveCount,omitempty" xml:"liveCount"`

	ThirdPartyHardCount *ResDetailDto `json:"thirdPartyHardCount,omitempty" xml:"thirdPartyHardCount"`

	HwVisionCount *ResDetailDto `json:"hwVisionCount,omitempty" xml:"hwVisionCount"`

	IdeaHubCount *ResDetailDto `json:"ideaHubCount,omitempty" xml:"ideaHubCount"`

	// 查询云会议室套餐包分配数量结果。
	Vmr *[]QueryVmrPkgResResultDto `json:"vmr,omitempty" xml:"vmr"`

	// 在创建企业的时候设置的pstn权限开关
	EnablePstn *bool `json:"enablePstn,omitempty" xml:"enablePstn"`

	// 在创建企业的时候设置的短信权限开关
	EnableSMS *bool `json:"enableSMS,omitempty" xml:"enableSMS"`

	// 企业是否开启混合云模式
	EnableHybridCloud *bool `json:"enableHybridCloud,omitempty" xml:"enableHybridCloud"`

	// 是否开启云盘
	EnableCloudDisk *bool `json:"enableCloudDisk,omitempty" xml:"enableCloudDisk"`

	// 是否开启UC功能
	EnableUc *bool `json:"enableUc,omitempty" xml:"enableUc"`

	// 是否开启Ai会议纪要
	EnableAiMinutes *bool `json:"enableAiMinutes,omitempty" xml:"enableAiMinutes"`

	// 单会议并发呼叫数
	SingleConfCallCount *int32 `json:"singleConfCallCount,omitempty" xml:"singleConfCallCount"`

	// 会议时长
	ConfLength     *int32 `json:"confLength,omitempty" xml:"confLength"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCorpResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpResourceResponse", string(data)}, " ")
}
