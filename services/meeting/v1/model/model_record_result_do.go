package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 录制信息
type RecordResultDo struct {

	// 会议UUID。
	ConfUUID *string `json:"confUUID,omitempty"`

	// 会议ID。
	ConfID *string `json:"confID,omitempty"`

	// 点播地址。
	Url *[]string `json:"url,omitempty"`

	// 录制时长（单位秒）。
	RcdTime *int32 `json:"rcdTime,omitempty"`

	// 录制文件大小（MB）。
	RcdSize *int32 `json:"rcdSize,omitempty"`

	// 会议主题。
	Subject *string `json:"subject,omitempty"`

	// 会议预订者。
	ScheduserName *string `json:"scheduserName,omitempty"`

	// 会议开始时间。
	StartTime *string `json:"startTime,omitempty"`

	// 录制文件是否转码完成。
	IsDecodeFinish *bool `json:"isDecodeFinish,omitempty"`

	// 录制文件预计转码完成时间。
	DecodeEndTime *int64 `json:"decodeEndTime,omitempty"`

	// 录播文件是否可观看。
	Available *bool `json:"available,omitempty"`

	// * 录播鉴权方式，在录播类型为:录播、直播+录播时有效 * 0： 可通过链接观看/下载 * 1： 企业用户可观看/下载 * 2： 与会者可观看/下载
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`
}

func (o RecordResultDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordResultDo struct{}"
	}

	return strings.Join([]string{"RecordResultDo", string(data)}, " ")
}
