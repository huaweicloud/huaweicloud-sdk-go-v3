package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移错误信息
type MigrationErrors struct {

	// 保存错误信息的json字符串
	ErrorJson *string `json:"error_json,omitempty" xml:"error_json"`

	// 主机名称（从用户系统获取，可能为空）
	HostName *string `json:"host_name,omitempty" xml:"host_name"`

	// 源端在主机迁移服务中的名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 源端服务器id
	SourceId *string `json:"source_id,omitempty" xml:"source_id"`

	// 源端服务器的ip
	SourceIp *string `json:"source_ip,omitempty" xml:"source_ip"`

	// 目的端服务器的ip
	TargetIp *string `json:"target_ip,omitempty" xml:"target_ip"`
}

func (o MigrationErrors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationErrors struct{}"
	}

	return strings.Join([]string{"MigrationErrors", string(data)}, " ")
}
