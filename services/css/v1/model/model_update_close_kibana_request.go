package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCloseKibanaRequest struct {
	// 指定待关闭的集群ID。

	ClusterId string `json:"cluster_id"`
}

func (o UpdateCloseKibanaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloseKibanaRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloseKibanaRequest", string(data)}, " ")
}
