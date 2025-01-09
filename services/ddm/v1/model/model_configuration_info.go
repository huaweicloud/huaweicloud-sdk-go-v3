package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationInfo 参数模板信息。
type ConfigurationInfo struct {

	// 参数模板ID。
	Id string `json:"id"`

	// 参数模板名称。
	Name string `json:"name"`

	// 参数模板描述。
	Description *string `json:"description,omitempty"`

	// 数据库类型。
	DatastoreName string `json:"datastore_name"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created string `json:"created"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated string `json:"updated"`

	// 是否是用户自定义参数模板：  false，表示为系统默认参数模板。 true，表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined"`
}

func (o ConfigurationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationInfo struct{}"
	}

	return strings.Join([]string{"ConfigurationInfo", string(data)}, " ")
}
