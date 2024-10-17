package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRdsDatabaseRequest Request Object
type AddRdsDatabaseRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *RdsDbRequest `json:"body,omitempty"`
}

func (o AddRdsDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRdsDatabaseRequest struct{}"
	}

	return strings.Join([]string{"AddRdsDatabaseRequest", string(data)}, " ")
}
