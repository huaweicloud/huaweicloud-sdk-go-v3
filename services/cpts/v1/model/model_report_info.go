package model

import (
	"encoding/json"

	"strings"
)

type ReportInfo struct {
	Brokens *ReportbrokensInfo `json:"brokens,omitempty"`

	Details *ReportdetailsInfo `json:"details,omitempty"`

	Outline *ReportoutlineInfo `json:"outline,omitempty"`
	// rtproportion

	Rtproportion *string `json:"rtproportion,omitempty"`

	TaskInfo *ReportTaskInfo `json:"taskInfo,omitempty"`
}

func (o ReportInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReportInfo struct{}"
	}

	return strings.Join([]string{"ReportInfo", string(data)}, " ")
}
