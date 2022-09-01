package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryRunInfo struct {

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// run_id
	RunId *float64 `json:"run_id,omitempty" xml:"run_id"`

	// run_type
	RunType *float64 `json:"run_type,omitempty" xml:"run_type"`

	// start_time
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// continue_time
	ContinueTime *float64 `json:"continue_time,omitempty" xml:"continue_time"`

	// temp_names
	TempNames *[]TempName `json:"temp_names,omitempty" xml:"temp_names"`
}

func (o HistoryRunInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryRunInfo struct{}"
	}

	return strings.Join([]string{"HistoryRunInfo", string(data)}, " ")
}
