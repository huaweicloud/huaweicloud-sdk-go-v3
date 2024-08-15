package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdmQueryClusterInstanceDetailFlavor 节点的虚拟机规格，请参见flavor参数说明。
type CdmQueryClusterInstanceDetailFlavor struct {

	// 节点虚拟机的规格ID。
	Id *string `json:"id,omitempty"`

	// 链接信息
	Links *[]ClusterLinks `json:"links,omitempty"`
}

func (o CdmQueryClusterInstanceDetailFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmQueryClusterInstanceDetailFlavor struct{}"
	}

	return strings.Join([]string{"CdmQueryClusterInstanceDetailFlavor", string(data)}, " ")
}
