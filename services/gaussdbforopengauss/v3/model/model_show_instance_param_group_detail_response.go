package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceParamGroupDetailResponse Response Object
type ShowInstanceParamGroupDetailResponse struct {

	// 引擎版本。
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// 引擎名称。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// 创建时间，格式为\"yyyy-MM-dd HH:mm:ss\"。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为\"yyyy-MM-ddHH:mm:ss\"。
	Updated *string `json:"updated,omitempty"`

	// 参数对象，用户基于默认参数模板自定义的参数配置。
	ConfigurationParameters *[]ConfigurationParameter `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                       `json:"-"`
}

func (o ShowInstanceParamGroupDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceParamGroupDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamGroupDetailResponse", string(data)}, " ")
}
