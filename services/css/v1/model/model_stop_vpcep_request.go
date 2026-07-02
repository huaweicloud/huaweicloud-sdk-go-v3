package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopVpcepRequest Request Object
type StopVpcepRequest struct {

	// 指定待关闭终端节点的集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o StopVpcepRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopVpcepRequest struct{}"
	}

	return strings.Join([]string{"StopVpcepRequest", string(data)}, " ")
}
