package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateUnbindPublicRequest struct {

	// 指定查询集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o UpdateUnbindPublicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUnbindPublicRequest struct{}"
	}

	return strings.Join([]string{"UpdateUnbindPublicRequest", string(data)}, " ")
}
