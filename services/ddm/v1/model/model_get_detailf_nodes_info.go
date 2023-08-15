package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDetailfNodesInfo nodes参数说明。
type GetDetailfNodesInfo struct {

	// DDM实例节点状态。
	Status string `json:"status"`

	// DDM实例节点port。
	Port string `json:"port"`

	// DDM实例节点IP。
	Ip string `json:"ip"`
}

func (o GetDetailfNodesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDetailfNodesInfo struct{}"
	}

	return strings.Join([]string{"GetDetailfNodesInfo", string(data)}, " ")
}
