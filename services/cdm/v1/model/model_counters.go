package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Counters struct {
	OrgApacheSqoopSubmissionCounterSqoopCounters *Counter `json:"org.apache.sqoop.submission.counter.SqoopCounters"`
}

func (o Counters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Counters struct{}"
	}

	return strings.Join([]string{"Counters", string(data)}, " ")
}
