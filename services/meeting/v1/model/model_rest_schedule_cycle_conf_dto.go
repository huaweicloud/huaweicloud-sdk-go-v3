package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestScheduleCycleConfDto struct {

	// 会议ID，长度限制为不超过32个字符
	CycleSubConfID string `json:"cycleSubConfID"`

	// 会议的媒体类型 由1个或多个枚举String组成，多个枚举时，每个枚举值之间通过”,”逗号分隔，枚举值如下： “Voice”：语音 “Video”：标清视频 “HDVideo”：高清视频（与Video互斥，如果同时选择Video、HDVideo，则系统默认选择Video） “Telepresence”：智真(与HDVideo、Video互斥，如果同时选择，系统使用Telepresence)—暂不支持 “Data”：多媒体（AS会根据系统配置决定是否自动添加Data）
	MediaTypes string `json:"mediaTypes"`

	// 会议开始时间，使用UTC时间 预定创建会议时，如果没有指定开始时间，或填空串，则表示会议马上开始 格式：YYYY-MM-DD HH:MM
	StartTime string `json:"startTime"`

	// 会议持续时长，单位分钟，最长1440，最短15
	Length int32 `json:"length"`

	// 会议是否自动启动录制，在录播类型为:录播、直播+录播时有效。 1 :true：自动启动录制 0 :false：不自动启动录制
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`

	ConfConfigInfo *CycleSubConfConfigDto `json:"confConfigInfo,omitempty"`

	// 录播鉴权方式，在录播类型为:录播、直播+录播时有效。 0为老的鉴权方式，url中携带token鉴权，1为企业内会议用户鉴权，2为会议内会议用户鉴权
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`

	// 会议描述，长度限制为200个字符
	Description *string `json:"description,omitempty"`
}

func (o RestScheduleCycleConfDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestScheduleCycleConfDto struct{}"
	}

	return strings.Join([]string{"RestScheduleCycleConfDto", string(data)}, " ")
}
