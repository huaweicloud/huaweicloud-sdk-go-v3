package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmStartClusterReq struct {
	// 集群启动操作，定义集群启动标识，为空对象

	Start *interface{} `json:"start"`
}

func (o CdmStartClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmStartClusterReq struct{}"
	}

	return strings.Join([]string{"CdmStartClusterReq", string(data)}, " ")
}
