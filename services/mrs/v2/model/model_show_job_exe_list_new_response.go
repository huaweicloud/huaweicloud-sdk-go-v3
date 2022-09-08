package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobExeListNewResponse struct {

	// 作业列表。
	JobList        *[]JobQueryBean `json:"job_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowJobExeListNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExeListNewResponse struct{}"
	}

	return strings.Join([]string{"ShowJobExeListNewResponse", string(data)}, " ")
}
