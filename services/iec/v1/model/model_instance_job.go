package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceJob struct {
	JobId *string `json:"job_id,omitempty"`

	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o InstanceJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceJob struct{}"
	}

	return strings.Join([]string{"InstanceJob", string(data)}, " ")
}
