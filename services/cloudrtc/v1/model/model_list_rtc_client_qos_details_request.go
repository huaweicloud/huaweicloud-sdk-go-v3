package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListRtcClientQosDetailsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 域名
	Domain *string `json:"domain,omitempty"`

	// 应用id
	AppId string `json:"app_id"`

	// 房间ID
	RoomId string `json:"room_id"`

	// 发送端用户
	UserId *string `json:"user_id,omitempty"`

	// 需查询接收端用户id
	PeerId *string `json:"peer_id,omitempty"`

	// 流号
	StreamId *string `json:"stream_id,omitempty"`

	// 判断上下行数据
	Direction *string `json:"direction,omitempty"`

	// 需查询的指标，填all则返回所有指标。多个指标使用','分割 - appcpu：端侧APP CPU使用率（appCpu） - syscpu：端侧系统 CPU使用率（deviceCpu） - abit：端侧音频码率kpbs（bitrate） - vbit：端侧视频码率kbps（bitRate） - dbit：端侧辅流码率kbps（bitRate） - vfps：端侧视频帧率fps（actFrameRate） - dfps：端侧辅流帧率fps（actFrameRate） - vblock：端侧视频卡顿率（统计大于等于600ms视频卡顿） - dblock：端侧辅流卡顿率（统计大于等于600ms辅流卡顿） - aloss：端侧音频丢包率（pktLoss） - vloss：端侧视频丢包率（pktLoss） - dloss：端侧辅流丢包率（pktLoss） - vwidth：端侧视频分辨率宽（actPicW） - vheight：端侧视频分辨率高（actPicH） - dwidth：端侧辅流分辨率宽（actPicW） - dheight：端侧辅流分辨率高（actPicH） - ajitter：端侧音频抖动率（jitter） - artt：端侧音频时延（rtt） - vjitter：端侧视频抖动率（jitter） - vrtt：端侧视频时延（rtt） - djitter：端侧辅流抖动率（jitter） - drtt：端侧辅流时延（rtt）
	Mid ListRtcClientQosDetailsRequestMid `json:"mid"`

	// 查询起始时间。UTC时间，格式：yyyy-mm-ddThh:mm:ssZ，如2020-04-23T06:00:00Z
	StartTime string `json:"start_time"`

	// 查询结束时间。UTC时间，格式：yyyy-mm-ddThh:mm:ssZ，如2020-04-23T07:00:00Z
	EndTime string `json:"end_time"`

	// 查询的时间类型取值：stime 数据库打点时间，不填默认ctime查询
	TimeType *string `json:"time_type,omitempty"`

	// 查询结果限制
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRtcClientQosDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcClientQosDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListRtcClientQosDetailsRequest", string(data)}, " ")
}

type ListRtcClientQosDetailsRequestMid struct {
	value string
}

type ListRtcClientQosDetailsRequestMidEnum struct {
	APPCPU  ListRtcClientQosDetailsRequestMid
	SYSCPU  ListRtcClientQosDetailsRequestMid
	ABIT    ListRtcClientQosDetailsRequestMid
	ABLOCK  ListRtcClientQosDetailsRequestMid
	VBIT    ListRtcClientQosDetailsRequestMid
	DBIT    ListRtcClientQosDetailsRequestMid
	VFPS    ListRtcClientQosDetailsRequestMid
	DFPS    ListRtcClientQosDetailsRequestMid
	VBLOCK  ListRtcClientQosDetailsRequestMid
	DBLOCK  ListRtcClientQosDetailsRequestMid
	ALOSS   ListRtcClientQosDetailsRequestMid
	VLOSS   ListRtcClientQosDetailsRequestMid
	DLOSS   ListRtcClientQosDetailsRequestMid
	VWIDTH  ListRtcClientQosDetailsRequestMid
	VHEIGHT ListRtcClientQosDetailsRequestMid
	DWIDTH  ListRtcClientQosDetailsRequestMid
	DHEIGHT ListRtcClientQosDetailsRequestMid
	AJITTER ListRtcClientQosDetailsRequestMid
	ARTT    ListRtcClientQosDetailsRequestMid
	VJITTER ListRtcClientQosDetailsRequestMid
	VRTT    ListRtcClientQosDetailsRequestMid
	DJITTER ListRtcClientQosDetailsRequestMid
	DRTT    ListRtcClientQosDetailsRequestMid
}

func GetListRtcClientQosDetailsRequestMidEnum() ListRtcClientQosDetailsRequestMidEnum {
	return ListRtcClientQosDetailsRequestMidEnum{
		APPCPU: ListRtcClientQosDetailsRequestMid{
			value: "appcpu",
		},
		SYSCPU: ListRtcClientQosDetailsRequestMid{
			value: "syscpu",
		},
		ABIT: ListRtcClientQosDetailsRequestMid{
			value: "abit",
		},
		ABLOCK: ListRtcClientQosDetailsRequestMid{
			value: "ablock",
		},
		VBIT: ListRtcClientQosDetailsRequestMid{
			value: "vbit",
		},
		DBIT: ListRtcClientQosDetailsRequestMid{
			value: "dbit",
		},
		VFPS: ListRtcClientQosDetailsRequestMid{
			value: "vfps",
		},
		DFPS: ListRtcClientQosDetailsRequestMid{
			value: "dfps",
		},
		VBLOCK: ListRtcClientQosDetailsRequestMid{
			value: "vblock",
		},
		DBLOCK: ListRtcClientQosDetailsRequestMid{
			value: "dblock",
		},
		ALOSS: ListRtcClientQosDetailsRequestMid{
			value: "aloss",
		},
		VLOSS: ListRtcClientQosDetailsRequestMid{
			value: "vloss",
		},
		DLOSS: ListRtcClientQosDetailsRequestMid{
			value: "dloss",
		},
		VWIDTH: ListRtcClientQosDetailsRequestMid{
			value: "vwidth",
		},
		VHEIGHT: ListRtcClientQosDetailsRequestMid{
			value: "vheight",
		},
		DWIDTH: ListRtcClientQosDetailsRequestMid{
			value: "dwidth",
		},
		DHEIGHT: ListRtcClientQosDetailsRequestMid{
			value: "dheight",
		},
		AJITTER: ListRtcClientQosDetailsRequestMid{
			value: "ajitter",
		},
		ARTT: ListRtcClientQosDetailsRequestMid{
			value: "artt",
		},
		VJITTER: ListRtcClientQosDetailsRequestMid{
			value: "vjitter",
		},
		VRTT: ListRtcClientQosDetailsRequestMid{
			value: "vrtt",
		},
		DJITTER: ListRtcClientQosDetailsRequestMid{
			value: "djitter",
		},
		DRTT: ListRtcClientQosDetailsRequestMid{
			value: "drtt",
		},
	}
}

func (c ListRtcClientQosDetailsRequestMid) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRtcClientQosDetailsRequestMid) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
