package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAvailableRdsList。
type QueryAvailableRdsList struct {

	// 数据库实例 ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 数据库实例所在租户在某一region下的project ID。
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`

	// 数据库实例状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 数据库实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 数据库实例引擎名称。
	EngineName *string `json:"engineName,omitempty" xml:"engineName"`

	// 数据库实例引擎版本。
	EngineSoftwareVersion *string `json:"engineSoftwareVersion,omitempty" xml:"engineSoftwareVersion"`

	// 数据库实例内网连接地址。
	PrivateIp *string `json:"privateIp,omitempty" xml:"privateIp"`

	// 数据库实例类型（主备或单机）。
	Mode *string `json:"mode,omitempty" xml:"mode"`

	// 数据库实例端口。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 可用区。
	AzCode *string `json:"azCode,omitempty" xml:"azCode"`

	// 时区。
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone"`
}

func (o QueryAvailableRdsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAvailableRdsList struct{}"
	}

	return strings.Join([]string{"QueryAvailableRdsList", string(data)}, " ")
}
