package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProvisionedThroughput struct {

	// 预置表级读请求单元数
	Rcu int64 `bson:"rcu"`

	// 预置表级写请求单元数
	Wcu int64 `bson:"wcu"`
}

func (o ProvisionedThroughput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProvisionedThroughput struct{}"
	}

	return strings.Join([]string{"ProvisionedThroughput", string(data)}, " ")
}
