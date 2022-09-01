package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfigurationsResult struct {

	// 参数模板ID。
	Id string `json:"id" xml:"id"`

	// 参数模板名称。
	Name string `json:"name" xml:"name"`

	// 参数模板描述。
	Description string `json:"description" xml:"description"`

	// 数据库版本。
	DatastoreVersion string `json:"datastore_version" xml:"datastore_version"`

	// 数据库类型。
	DatastoreName string `json:"datastore_name" xml:"datastore_name"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created string `json:"created" xml:"created"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated string `json:"updated" xml:"updated"`

	// 是否是用户自定义参数模板。 - false表示为默认参数模板。 - true表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined" xml:"user_defined"`
}

func (o ListConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResult", string(data)}, " ")
}
