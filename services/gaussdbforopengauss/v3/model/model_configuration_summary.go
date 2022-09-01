package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数模板信息。
type ConfigurationSummary struct {

	// 参数组ID。
	Id string `json:"id" xml:"id"`

	// 参数组名称。
	Name string `json:"name" xml:"name"`

	// 参数组描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 引擎版本。
	DatastoreVersion string `json:"datastore_version" xml:"datastore_version"`

	// 引擎名称。
	DatastoreName string `json:"datastore_name" xml:"datastore_name"`

	// 数据库部署模式。
	HaMode string `json:"ha_mode" xml:"ha_mode"`

	// 创建时间，格式为\"yyyy-MM-dd HH:mm:ss\"。
	Created string `json:"created" xml:"created"`

	// 更新时间，格式为\"yyyy-MM-dd HH:mm:ss\"。
	Updated string `json:"updated" xml:"updated"`

	// 是否是用户自定义参数模板：  - false，表示为系统默认参数模板。 - true，表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined" xml:"user_defined"`
}

func (o ConfigurationSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationSummary struct{}"
	}

	return strings.Join([]string{"ConfigurationSummary", string(data)}, " ")
}
