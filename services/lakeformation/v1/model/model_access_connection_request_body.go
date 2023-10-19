package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConnectionRequestBody 创建接入连接的请求信息
type AccessConnectionRequestBody struct {

	// 虚拟私有云终端节点ID。在 接入管理-创建客户端-前往VPC创建-VPC终端节点 创建和查看。
	VpcepId *string `json:"vpcep_id,omitempty"`

	// 终端节点服务名称。最大长度为64个字符。
	VpcepServiceName *string `json:"vpcep_service_name,omitempty"`

	// 接入域名，通过IP接入访问Lakeformation API时，需在请求头中添加HOST参数并传入该域名。
	Domain *string `json:"domain,omitempty"`
}

func (o AccessConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"AccessConnectionRequestBody", string(data)}, " ")
}
