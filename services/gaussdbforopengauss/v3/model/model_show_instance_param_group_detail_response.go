package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceParamGroupDetailResponse Response Object
type ShowInstanceParamGroupDetailResponse struct {

	// **参数解释**: 实例对应参数组ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 引擎版本。 **取值范围**: 不涉及。
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// **参数解释**: 引擎名称。 **取值范围**: 不涉及。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// **参数解释**: 创建时间，格式为\"yyyy-MM-dd HH:mm:ss\"。 **取值范围**: 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**: 更新时间，格式为\"yyyy-MM-dd HH:mm:ss\"。 **取值范围**: 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**: 参数对象，用户基于默认参数模板自定义的参数配置，具体请参考ConfigurationParameterResult。
	ConfigurationParameters *[]ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                             `json:"-"`
}

func (o ShowInstanceParamGroupDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceParamGroupDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamGroupDetailResponse", string(data)}, " ")
}
