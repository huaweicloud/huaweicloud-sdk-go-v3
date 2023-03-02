package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSpecIssueStayTimesResponseBodyData struct {

	// 工作项id字符串
	Id *string `json:"id,omitempty"`

	// 停留时间（单位：秒）
	StayTime *int64 `json:"stay_time,omitempty"`
}

func (o ListSpecIssueStayTimesResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecIssueStayTimesResponseBodyData struct{}"
	}

	return strings.Join([]string{"ListSpecIssueStayTimesResponseBodyData", string(data)}, " ")
}
