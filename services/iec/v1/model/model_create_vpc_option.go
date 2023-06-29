package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcOption 虚拟私有云的请求数据对象。
type CreateVpcOption struct {

	// 虚拟私有云名称  取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）  约束：同一个帐号下的名称不能重复
	Name string `json:"name"`

	// 虚拟私有云下可用子网的范围。  约束： SYSTEM模式，cidr取值范围：10.0.0.0/8~10.255.0.0/16或者172.16.0.0/12 ~ 172.31.0.0/16或者192.168.0.0/16 。 [CUSTOMER模式，cidr的取值范围：10.0.0.0/8~10.255.255.0/24或者172.16.0.0/12 ~ 172.32.255.0/24或者192.168.0.0/16~192.168.255.0/24。](tag:internal)
	Cidr string `json:"cidr"`

	// 虚拟私有云的模式，支持的取值范围如下：  SYSTEM：该类型网络，系统会自动按照实际需要创建足够的子网。 [CUSTOMER：该类型网络，用户需要完全按照自己站点的需要，去申请足够的子网。](tag:internal)
	Mode string `json:"mode"`
}

func (o CreateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcOption struct{}"
	}

	return strings.Join([]string{"CreateVpcOption", string(data)}, " ")
}
