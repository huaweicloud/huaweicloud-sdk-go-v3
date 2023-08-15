package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HbaseClusterActionReq struct {

	// 该请求参数内无其他内容，但是需要该参数作为重启集群入参，示例看下述所示
	Restart *interface{} `json:"restart"`
}

func (o HbaseClusterActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HbaseClusterActionReq struct{}"
	}

	return strings.Join([]string{"HbaseClusterActionReq", string(data)}, " ")
}
