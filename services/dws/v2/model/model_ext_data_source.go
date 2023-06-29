package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtDataSource 数据源信息
type ExtDataSource struct {

	// id。
	Id *string `json:"id,omitempty"`

	// 数据源名称。
	Name *string `json:"name,omitempty"`

	// 类型。
	Type *string `json:"type,omitempty"`

	// 数据库。
	ConnectInfo *string `json:"connect_info,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 版本。
	Version *string `json:"version,omitempty"`

	// 配置状态。
	ConfigureStatus *string `json:"configure_status,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 数据源id。
	DataSourceId *string `json:"data_source_id,omitempty"`

	// 创建时间。
	Created *string `json:"created,omitempty"`

	// 更新时间。
	Updated *string `json:"updated,omitempty"`

	// 数据源更新时间。
	DataSourceUpdated *string `json:"data_source_updated,omitempty"`

	// 扩展信息。
	ExtendProperties *interface{} `json:"extend_properties,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o ExtDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtDataSource struct{}"
	}

	return strings.Join([]string{"ExtDataSource", string(data)}, " ")
}
