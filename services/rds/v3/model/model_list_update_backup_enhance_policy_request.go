package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdateBackupEnhancePolicyRequest Request Object
type ListUpdateBackupEnhancePolicyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 用户当前时区，例：plus08
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ListUpdateBackupEnhancePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdateBackupEnhancePolicyRequest struct{}"
	}

	return strings.Join([]string{"ListUpdateBackupEnhancePolicyRequest", string(data)}, " ")
}
