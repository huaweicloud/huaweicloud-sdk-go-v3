package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataSyncStatistics struct {
	AbnormalCount *int32 `json:"abnormal_count,omitempty"`

	RunningCount *int32 `json:"running_count,omitempty"`

	TypeTwoWayCount *int32 `json:"type_two_way_count,omitempty"`

	TypeUnidirectionalCount *int32 `json:"type_unidirectional_count,omitempty"`
}

func (o DataSyncStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSyncStatistics struct{}"
	}

	return strings.Join([]string{"DataSyncStatistics", string(data)}, " ")
}
