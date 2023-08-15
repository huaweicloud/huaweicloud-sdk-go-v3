package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterRecoveryProgress 集群容灾进度详情
type ClusterRecoveryProgress struct {

	// key
	Key *string `json:"key,omitempty"`

	// action_type
	ActionType *string `json:"action_type,omitempty"`

	// unrestore_keys
	UnrestoreKeys *string `json:"unrestore_keys,omitempty"`

	// action_start_time
	ActionStartTime *string `json:"action_start_time,omitempty"`

	// action_end_time
	ActionEndTime *string `json:"action_end_time,omitempty"`
}

func (o ClusterRecoveryProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRecoveryProgress struct{}"
	}

	return strings.Join([]string{"ClusterRecoveryProgress", string(data)}, " ")
}
