package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJarPackageHostInfoRequest Request Object
type ListJarPackageHostInfoRequest struct {

	// 租户企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 文件名称
	FileName string `json:"file_name"`

	// 类别，包含如下:   - host : 主机   - container : 容器
	Category *string `json:"category,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器IP
	HostIp *string `json:"host_ip,omitempty"`

	// 默认10
	Limit *int32 `json:"limit,omitempty"`

	// 默认是0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListJarPackageHostInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageHostInfoRequest struct{}"
	}

	return strings.Join([]string{"ListJarPackageHostInfoRequest", string(data)}, " ")
}
