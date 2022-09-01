package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 选看消息体。
type RestChairViewReqBody struct {

	// 主持人观看的画面类型。 - 0: 主持人轮询。 - 1: 主持人观看多画面。 - 2: 主持人选看会场。
	ViewType int32 `json:"viewType" xml:"viewType"`

	// 被主持人选看的会场。 当为主持人选看会场时为必填参数。
	ParticipantID *string `json:"participantID,omitempty" xml:"participantID"`

	// 主持人轮询时，必填字段。 表示轮询间隔，单位：秒。 范围:[10-120]，默认值：30
	SwitchTime *int32 `json:"switchTime,omitempty" xml:"switchTime"`

	// 主持人轮询时，必填字段。 - 0: 停止轮询。 - 1: 启动轮询。
	Status *int32 `json:"status,omitempty" xml:"status"`
}

func (o RestChairViewReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestChairViewReqBody struct{}"
	}

	return strings.Join([]string{"RestChairViewReqBody", string(data)}, " ")
}
