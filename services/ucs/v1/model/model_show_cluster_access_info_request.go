package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterAccessInfoRequest Request Object
type ShowClusterAccessInfoRequest struct {

	// 集群ID
	Clusterid string `json:"clusterid"`

	// VPC终端节点的IP地址。私网接入的集群必填，且必须是打通线下集群的VPC终端节点。  创建VPC终端节点及查询IP地址的方法请参见[创建终端节点](https://support.huaweicloud.com/usermanual-ucs/ucs_01_0336.html#section2)。
	Vpcendpoint *string `json:"vpcendpoint,omitempty"`

	// 接入region，私网接入的集群必填
	Region *string `json:"region,omitempty"`
}

func (o ShowClusterAccessInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterAccessInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterAccessInfoRequest", string(data)}, " ")
}
