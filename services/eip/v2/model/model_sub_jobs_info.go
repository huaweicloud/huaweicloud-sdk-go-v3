package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubJobsInfo struct {

	// 子job信息，类型与主job一致
	SubJobs *[]interface{} `json:"sub_jobs,omitempty"`
}

func (o SubJobsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobsInfo struct{}"
	}

	return strings.Join([]string{"SubJobsInfo", string(data)}, " ")
}
