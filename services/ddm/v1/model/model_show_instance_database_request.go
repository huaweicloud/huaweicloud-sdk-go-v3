package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDatabaseRequest Request Object
type ShowInstanceDatabaseRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DatabaseName string `json:"database_name"`
}

func (o ShowInstanceDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDatabaseRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDatabaseRequest", string(data)}, " ")
}
