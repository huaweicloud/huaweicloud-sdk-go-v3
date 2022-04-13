package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAz2MigrateRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ListAz2MigrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAz2MigrateRequest struct{}"
	}

	return strings.Join([]string{"ListAz2MigrateRequest", string(data)}, " ")
}
