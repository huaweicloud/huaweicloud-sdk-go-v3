package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新测试类型服务信息
type ServiceRequestBody struct {

	// 测试类型名称，用于界面显示，不能使用当前保留名，长度小于等于16位字符
	ServiceName string `json:"service_name"`

	// server_host是由用户提供的域名。 我们会通过此域名进行接口调用，请以https/http开头,长度小于等于128位字符。 TestHub将会通过此域名下的接口，保证云测数据与用户系统数据的一致性。
	ServerHost string `json:"server_host"`
}

func (o ServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceRequestBody struct{}"
	}

	return strings.Join([]string{"ServiceRequestBody", string(data)}, " ")
}
