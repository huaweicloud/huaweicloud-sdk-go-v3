package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportInfo struct {
	Brokens *ReportbrokensInfo `json:"brokens,omitempty" xml:"brokens"`

	Details *ReportdetailsInfo `json:"details,omitempty" xml:"details"`

	Outline *ReportoutlineInfo `json:"outline,omitempty" xml:"outline"`

	// 响应时间分布
	Rtproportion *string `json:"rtproportion,omitempty" xml:"rtproportion"`

	TaskInfo *ReportTaskInfo `json:"taskInfo,omitempty" xml:"taskInfo"`
}

func (o ReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportInfo struct{}"
	}

	return strings.Join([]string{"ReportInfo", string(data)}, " ")
}
