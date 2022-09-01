package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业资源信息
type QueryCorpVcResResultDto struct {

	// 虚拟会议室类型列表,最多支持8个，暂不限制
	VmrPkgList *[]QueryVmrPkgResResultDto `json:"vmrPkgList,omitempty" xml:"vmrPkgList"`

	// 1080P硬终端账户数
	Te1080pHardCount *int32 `json:"te1080pHardCount,omitempty" xml:"te1080pHardCount"`

	// 720P硬终端账户数
	Te720pHardCount *int32 `json:"te720pHardCount,omitempty" xml:"te720pHardCount"`

	// 软终端账户数
	TeSoftCount *int32 `json:"teSoftCount,omitempty" xml:"teSoftCount"`

	// 大屏软终端数量
	RoomCount *int32 `json:"roomCount,omitempty" xml:"roomCount"`

	// 录播存储空间 （单位：G）
	RecordCapability *int32 `json:"recordCapability,omitempty" xml:"recordCapability"`

	// 会议并发方数
	ConfCallCount *int32 `json:"confCallCount,omitempty" xml:"confCallCount"`

	// 推流并发数量
	LiveCount *int32 `json:"liveCount,omitempty" xml:"liveCount"`

	// 第三方硬终端接入数
	ThirdPartyHardCount *int32 `json:"thirdPartyHardCount,omitempty" xml:"thirdPartyHardCount"`

	// 智慧屏终端接入数
	HwVisionCount *int32 `json:"hwVisionCount,omitempty" xml:"hwVisionCount"`

	// ideahub终端接入数
	IdeaHubCount *int32 `json:"ideaHubCount,omitempty" xml:"ideaHubCount"`
}

func (o QueryCorpVcResResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpVcResResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpVcResResultDto", string(data)}, " ")
}
