package model

import (
	"encoding/json"

	"strings"
)

type Counters struct {
	OrgApacheSqoopSubmissionCounterSqoopCounters *Counter `json:"org.apache.sqoop.submission.counter.SqoopCounters"`
}

func (o Counters) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Counters struct{}"
	}

	return strings.Join([]string{"Counters", string(data)}, " ")
}
