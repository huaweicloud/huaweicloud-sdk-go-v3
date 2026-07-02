package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartVpcepRequest Request Object
type StartVpcepRequest struct {

	// 指定开启终端节点的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *StartVpecpReq `json:"body,omitempty"`
}

func (o StartVpcepRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartVpcepRequest struct{}"
	}

	return strings.Join([]string{"StartVpcepRequest", string(data)}, " ")
}
