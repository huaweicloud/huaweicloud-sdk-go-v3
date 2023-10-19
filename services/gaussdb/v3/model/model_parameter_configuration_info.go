package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterConfigurationInfo 配置信息。
type ParameterConfigurationInfo struct {

	// 数据库版本名称。
	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	// 数据库名称。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Updated *string `json:"updated,omitempty"`
}

func (o ParameterConfigurationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterConfigurationInfo struct{}"
	}

	return strings.Join([]string{"ParameterConfigurationInfo", string(data)}, " ")
}
