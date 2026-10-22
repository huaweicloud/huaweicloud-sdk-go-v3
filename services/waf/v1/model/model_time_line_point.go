package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeLinePoint 时间点统计数据
type TimeLinePoint struct {

	// 时间点（毫秒时间戳），代表以该时间点截止的5分钟内的统计数据
	Time int64 `json:"time"`

	// 数量
	Num int64 `json:"num"`
}

func (o TimeLinePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeLinePoint struct{}"
	}

	return strings.Join([]string{"TimeLinePoint", string(data)}, " ")
}
