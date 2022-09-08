package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CycleSubConf struct {

	// 子会议ID
	CycleSubConfID string `json:"cycleSubConfID"`

	// 会议ID，长度限制为不超过32个字符
	ConferenceID *string `json:"conferenceID,omitempty"`

	// 会议的媒体类型。 由1个或多个枚举String组成，多个枚举时，每个枚举值之间通过”,”逗号分隔，枚举值如下： “Voice”：语音 “Video”：标清视频 “HDVideo”：高清视频（与Video互斥，如果同时选择Video、HDVideo，则系统默认选择Video） “Telepresence”：智真(与HDVideo、Video互斥，如果同时选择，系统使用Telepresence)—暂不支持 “Data”：多媒体
	MediaType *string `json:"mediaType,omitempty"`

	// 会议起始时间(格式：YYYY-MM-DD HH:MM)
	StartTime *string `json:"startTime,omitempty"`

	// 会议结束时间(格式：YYYY-MM-DD HH:MM)
	EndTime *string `json:"endTime,omitempty"`

	// 是否自动开启录音
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`

	ConfConfigInfo *CycleSubConfConfigDto `json:"confConfigInfo,omitempty"`

	// 录播鉴权方式，在录播类型为:录播、直播+录播时有效。 0为老的鉴权方式，url中携带token鉴权，1为企业内会议用户鉴权，2为会议内会议用户鉴权
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`

	// 会议描述，长度限制为200个字符
	Description *string `json:"description,omitempty"`
}

func (o CycleSubConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CycleSubConf struct{}"
	}

	return strings.Join([]string{"CycleSubConf", string(data)}, " ")
}
