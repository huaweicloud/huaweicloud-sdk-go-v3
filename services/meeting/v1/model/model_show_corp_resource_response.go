package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCorpResourceResponse struct {
	Te1080pHardCount *ResDetailDto `json:"te1080pHardCount,omitempty"`

	Te720pHardCount *ResDetailDto `json:"te720pHardCount,omitempty"`

	TeSoftCount *ResDetailDto `json:"teSoftCount,omitempty"`

	RoomCount *ResDetailDto `json:"roomCount,omitempty"`

	RecordCapability *ResDetailDto `json:"recordCapability,omitempty"`

	ConfCallCount *ResDetailDto `json:"confCallCount,omitempty"`

	LiveCount *ResDetailDto `json:"liveCount,omitempty"`

	ThirdPartyHardCount *ResDetailDto `json:"thirdPartyHardCount,omitempty"`

	HwVisionCount *ResDetailDto `json:"hwVisionCount,omitempty"`

	IdeaHubCount *ResDetailDto `json:"ideaHubCount,omitempty"`

	// 查询云会议室套餐包分配数量结果。
	Vmr *[]QueryVmrPkgResResultDto `json:"vmr,omitempty"`

	// 在创建企业的时候设置的pstn权限开关。
	EnablePstn *bool `json:"enablePstn,omitempty"`

	// 企业是否通过短信形式发送会议通知。
	EnableSMS *bool `json:"enableSMS,omitempty"`

	// 企业是否开启混合云模式。
	EnableHybridCloud *bool `json:"enableHybridCloud,omitempty"`

	// 是否开启云盘。
	EnableCloudDisk *bool `json:"enableCloudDisk,omitempty"`

	// 是否开启UC功能。
	EnableUc *bool `json:"enableUc,omitempty"`

	// 是否开启Ai会议纪要。
	EnableAiMinutes *bool `json:"enableAiMinutes,omitempty"`

	// 单会议并发呼叫数。
	SingleConfCallCount *int32 `json:"singleConfCallCount,omitempty"`

	// 会议时长。
	ConfLength     *int32 `json:"confLength,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCorpResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpResourceResponse", string(data)}, " ")
}
