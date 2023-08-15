package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsBuckets
type ObsBuckets struct {

	//
	ObsBuckets []string `json:"obs_buckets"`
}

func (o ObsBuckets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBuckets struct{}"
	}

	return strings.Join([]string{"ObsBuckets", string(data)}, " ")
}
