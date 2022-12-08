package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数模板信息。
type ConfigurationSummary2 struct {

	// 参数组ID。
	Id *string `json:"id,omitempty"`

	// 参数组名称。
	Name *string `json:"name,omitempty"`

	// 参数组描述。
	Description *string `json:"description,omitempty"`

	Datastore *DatastoreResult `json:"datastore,omitempty"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量。
	Updated *string `json:"updated,omitempty"`
}

func (o ConfigurationSummary2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationSummary2 struct{}"
	}

	return strings.Join([]string{"ConfigurationSummary2", string(data)}, " ")
}
