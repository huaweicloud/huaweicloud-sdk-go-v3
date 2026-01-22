package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationParameterResponse Response Object
type ShowConfigurationParameterResponse struct {

	// **参数解释：** 参数模板ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 参数模板名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 数据库版本。 **取值范围：** 不涉及。
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// **参数解释：** 数据库类型。 **取值范围：** 不涉及。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// **参数解释：** 参数模板描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 创建时间，格式为“yyyy-MM-ddTHH:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量。 **取值范围：** 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释：** 更新时间，格式为“yyyy-MM-ddTHH:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量。 **取值范围：** 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释：** 参数对象，用户基于默认参数模板自定义的参数配置。 **取值范围：** 不涉及。
	Parameters     *[]ConfigurationParametersResult `json:"parameters,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowConfigurationParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationParameterResponse", string(data)}, " ")
}
