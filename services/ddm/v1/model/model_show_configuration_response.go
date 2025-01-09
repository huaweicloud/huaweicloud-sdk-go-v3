package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationResponse Response Object
type ShowConfigurationResponse struct {

	// 参数组ID。
	Id *string `json:"id,omitempty"`

	// 参数组名称。
	Name *string `json:"name,omitempty"`

	// 数据库名称。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated *string `json:"updated,omitempty"`

	// 参数对象，用户基于默认参数模板自定义的参数配置。
	ConfigurationParameters *[]ConfigurationParameter `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                       `json:"-"`
}

func (o ShowConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationResponse", string(data)}, " ")
}
