package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempRunningData struct {

	// content_method_url
	ContentMethodUrl *[]string `json:"content_method_url,omitempty" xml:"content_method_url"`

	// crawler_status
	CrawlerStatus *int32 `json:"crawler_status,omitempty" xml:"crawler_status"`

	// related_temp_running_id
	RelatedTempRunningId *int32 `json:"related_temp_running_id,omitempty" xml:"related_temp_running_id"`

	// task_run_info_id
	TaskRunInfoId *int32 `json:"task_run_info_id,omitempty" xml:"task_run_info_id"`

	// temp_id
	TempId *int32 `json:"temp_id,omitempty" xml:"temp_id"`

	// temp_name
	TempName *string `json:"temp_name,omitempty" xml:"temp_name"`

	// temp_running_status
	TempRunningStatus *int32 `json:"temp_running_status,omitempty" xml:"temp_running_status"`
}

func (o TempRunningData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempRunningData struct{}"
	}

	return strings.Join([]string{"TempRunningData", string(data)}, " ")
}
