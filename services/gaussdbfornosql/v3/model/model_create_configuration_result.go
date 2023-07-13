package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationResult 参数模板信息。
type CreateConfigurationResult struct {

	// 参数模板ID。
	Id string `json:"id"`

	// 参数模板名称。
	Name string `json:"name"`

	// 参数模板描述。
	Description *string `json:"description,omitempty"`

	// 数据库版本名称。
	DatastoreVersionName string `json:"datastore_version_name"`

	// 数据库名称。
	DatastoreName string `json:"datastore_name"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created string `json:"created"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated string `json:"updated"`
}

func (o CreateConfigurationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationResult struct{}"
	}

	return strings.Join([]string{"CreateConfigurationResult", string(data)}, " ")
}
