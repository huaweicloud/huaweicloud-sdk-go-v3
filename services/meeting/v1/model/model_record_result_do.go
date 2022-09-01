package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 录制信息
type RecordResultDo struct {

	// 会议UUID。
	ConfUUID *string `json:"confUUID,omitempty" xml:"confUUID"`

	// 会议ID。
	ConfID *string `json:"confID,omitempty" xml:"confID"`

	// 点播地址。
	Url *[]string `json:"url,omitempty" xml:"url"`

	// 录制时长（单位秒）。
	RcdTime *int32 `json:"rcdTime,omitempty" xml:"rcdTime"`

	// 录制文件大小（MB）。
	RcdSize *int32 `json:"rcdSize,omitempty" xml:"rcdSize"`

	// 会议主题。
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 会议预订者。
	ScheduserName *string `json:"scheduserName,omitempty" xml:"scheduserName"`

	// 会议开始时间。
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 录制文件是否转码完成。
	IsDecodeFinish *bool `json:"isDecodeFinish,omitempty" xml:"isDecodeFinish"`

	// 录制文件预计转码完成时间。
	DecodeEndTime *int64 `json:"decodeEndTime,omitempty" xml:"decodeEndTime"`

	// 录播文件是否可观看。
	Available *bool `json:"available,omitempty" xml:"available"`

	// * 录播鉴权方式，在录播类型为:录播、直播+录播时有效 * 0： 可通过链接观看/下载 * 1： 企业用户可观看/下载 * 2： 与会者可观看/下载
	RecordAuthType *int32 `json:"recordAuthType,omitempty" xml:"recordAuthType"`
}

func (o RecordResultDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordResultDo struct{}"
	}

	return strings.Join([]string{"RecordResultDo", string(data)}, " ")
}
