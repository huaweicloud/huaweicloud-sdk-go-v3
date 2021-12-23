package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新测试类型服务信息
type ServiceRequestBody struct {
	// 测试类型名称，用于界面显示，不能使用当前保留名，长度小于等于16位字符

	ServiceName string `json:"service_name"`
	// 新测试类型服务域名，用于拼接调用接口，以https/http开头，长度小于等于128位字符

	ServerHost string `json:"server_host"`
}

func (o ServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceRequestBody struct{}"
	}

	return strings.Join([]string{"ServiceRequestBody", string(data)}, " ")
}
