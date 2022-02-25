package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVpcepConnectionRequest struct {
	// 指定待获取终端节点连接的集群ID。

	ClusterId string `json:"cluster_id"`
}

func (o ShowVpcepConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcepConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcepConnectionRequest", string(data)}, " ")
}
