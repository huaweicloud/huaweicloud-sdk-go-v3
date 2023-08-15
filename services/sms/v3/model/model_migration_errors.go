package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationErrors 迁移错误信息
type MigrationErrors struct {

	// 保存错误信息的json字符串
	ErrorJson *string `json:"error_json,omitempty"`

	// 主机名称（从用户系统获取，可能为空）
	HostName *string `json:"host_name,omitempty"`

	// 源端在主机迁移服务中的名称
	Name *string `json:"name,omitempty"`

	// 源端服务器ID
	SourceId *string `json:"source_id,omitempty"`

	// 源端服务器的ip
	SourceIp *string `json:"source_ip,omitempty"`

	// 目的端服务器的ip
	TargetIp *string `json:"target_ip,omitempty"`
}

func (o MigrationErrors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationErrors struct{}"
	}

	return strings.Join([]string{"MigrationErrors", string(data)}, " ")
}
