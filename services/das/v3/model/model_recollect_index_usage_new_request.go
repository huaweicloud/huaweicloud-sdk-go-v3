package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecollectIndexUsageNewRequest Request Object
type RecollectIndexUsageNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o RecollectIndexUsageNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecollectIndexUsageNewRequest struct{}"
	}

	return strings.Join([]string{"RecollectIndexUsageNewRequest", string(data)}, " ")
}
