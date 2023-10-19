package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConnectionInfo 接入连接信息
type AccessConnectionInfo struct {

	// 虚拟私有云终端节点ID。在 接入管理-创建客户端-前往VPC创建-VPC终端节点 创建和查看。
	VpcepId *string `json:"vpcep_id,omitempty"`

	// 接入IP
	Ip *string `json:"ip,omitempty"`

	// 拥有者
	Owner *string `json:"owner,omitempty"`

	// 接入域名，通过IP接入访问Lakeformation API时，需在请求头中添加HOST参数并传入该域名。
	Domain *string `json:"domain,omitempty"`
}

func (o AccessConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConnectionInfo struct{}"
	}

	return strings.Join([]string{"AccessConnectionInfo", string(data)}, " ")
}
