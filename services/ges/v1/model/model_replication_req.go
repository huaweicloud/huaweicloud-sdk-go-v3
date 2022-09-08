package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReplicationReq struct {

	// 新扩副本数量。
	Replication string `json:"replication"`
}

func (o ReplicationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationReq struct{}"
	}

	return strings.Join([]string{"ReplicationReq", string(data)}, " ")
}
