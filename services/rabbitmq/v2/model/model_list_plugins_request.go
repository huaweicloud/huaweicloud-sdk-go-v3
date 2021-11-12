package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPluginsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ListPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListPluginsRequest", string(data)}, " ")
}
