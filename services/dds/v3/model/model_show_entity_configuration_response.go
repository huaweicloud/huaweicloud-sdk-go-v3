package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEntityConfigurationResponse struct {

	// 数据库版本。
	DatastoreVersion *string `json:"datastore_version,omitempty" xml:"datastore_version"`

	// 数据库类型。
	DatastoreName *string `json:"datastore_name,omitempty" xml:"datastore_name"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created *string `json:"created,omitempty" xml:"created"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 参数对象，用户基于默认参数模板自定义的参数配置。
	Parameters     *[]EntityConfigurationParametersResult `json:"parameters,omitempty" xml:"parameters"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowEntityConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntityConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowEntityConfigurationResponse", string(data)}, " ")
}
