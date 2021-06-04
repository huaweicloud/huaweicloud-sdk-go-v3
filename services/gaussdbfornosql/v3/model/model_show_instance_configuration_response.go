package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceConfigurationResponse struct {
	// 数据库版本名称。

	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`
	// 数据库名称。

	DatastoreName *string `json:"datastore_name,omitempty"`
	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。

	Created *string `json:"created,omitempty"`
	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。

	Updated *string `json:"updated,omitempty"`
	// 参数对象，用户基于默认参数模板自定义的参数配置。

	ConfigurationParameters *[]ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                             `json:"-"`
}

func (o ShowInstanceConfigurationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationResponse", string(data)}, " ")
}
