package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 受配置检测影响的服务器信息
type SecurityCheckHostInfoResponseInfo struct {

	// 服务器ID
	HostId *string `json:"host_id,omitempty" xml:"host_id"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty" xml:"host_name"`

	// 服务器公网IP
	HostPublicIp *string `json:"host_public_ip,omitempty" xml:"host_public_ip"`

	// 服务器私网IP
	HostPrivateIp *string `json:"host_private_ip,omitempty" xml:"host_private_ip"`

	// 扫描时间
	ScanTime *int64 `json:"scan_time,omitempty" xml:"scan_time"`

	// 风险项数量
	FailedNum *int32 `json:"failed_num,omitempty" xml:"failed_num"`

	// 通过项数量
	PassedNum *int32 `json:"passed_num,omitempty" xml:"passed_num"`
}

func (o SecurityCheckHostInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckHostInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckHostInfoResponseInfo", string(data)}, " ")
}
