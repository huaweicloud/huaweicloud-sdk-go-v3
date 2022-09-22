package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业资源信息。
type QueryCorpVcResResultDto struct {

	// 云会议室类型列表。
	VmrPkgList *[]QueryVmrPkgResResultDto `json:"vmrPkgList,omitempty"`

	// 1080P硬终端接入帐号数量。
	Te1080pHardCount *int32 `json:"te1080pHardCount,omitempty"`

	// 720P硬终端接入帐号数量。
	Te720pHardCount *int32 `json:"te720pHardCount,omitempty"`

	// 软终端账户数量。
	TeSoftCount *int32 `json:"teSoftCount,omitempty"`

	// 电子白板（SmartRooms）接入帐号数量。
	RoomCount *int32 `json:"roomCount,omitempty"`

	// 录播存储空间 (单位:G)。
	RecordCapability *int32 `json:"recordCapability,omitempty"`

	// 会议并发方数量。
	ConfCallCount *int32 `json:"confCallCount,omitempty"`

	// 直播端口数量。
	LiveCount *int32 `json:"liveCount,omitempty"`

	// 第三方硬终端接入帐号数量。
	ThirdPartyHardCount *int32 `json:"thirdPartyHardCount,omitempty"`

	// 智慧屏终端接入帐号数量。
	HwVisionCount *int32 `json:"hwVisionCount,omitempty"`

	// IdeaHub终端接入帐号数量。
	IdeaHubCount *int32 `json:"ideaHubCount,omitempty"`
}

func (o QueryCorpVcResResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpVcResResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpVcResResultDto", string(data)}, " ")
}
