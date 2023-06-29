package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAvailableRdsList QueryAvailableRdsList。
type QueryAvailableRdsList struct {

	// 数据库实例 ID。
	Id *string `json:"id,omitempty"`

	// 数据库实例所在租户在某一region下的project ID。
	ProjectId *string `json:"projectId,omitempty"`

	// 数据库实例状态。
	Status *string `json:"status,omitempty"`

	// 数据库实例名称。
	Name *string `json:"name,omitempty"`

	// 数据库实例引擎名称。
	EngineName *string `json:"engineName,omitempty"`

	// 数据库实例引擎版本。
	EngineSoftwareVersion *string `json:"engineSoftwareVersion,omitempty"`

	// 数据库实例内网连接地址。
	PrivateIp *string `json:"privateIp,omitempty"`

	// 数据库实例类型（主备或单机）。
	Mode *string `json:"mode,omitempty"`

	// 数据库实例端口。
	Port *int32 `json:"port,omitempty"`

	// 可用区。
	AzCode *string `json:"azCode,omitempty"`

	// 时区。
	TimeZone *string `json:"timeZone,omitempty"`
}

func (o QueryAvailableRdsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAvailableRdsList struct{}"
	}

	return strings.Join([]string{"QueryAvailableRdsList", string(data)}, " ")
}
