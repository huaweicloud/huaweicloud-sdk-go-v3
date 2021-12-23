package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportTaskInfo struct {
	// vum

	Vum *int32 `json:"vum,omitempty"`
}

func (o ReportTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportTaskInfo struct{}"
	}

	return strings.Join([]string{"ReportTaskInfo", string(data)}, " ")
}
