package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConfigurationTemplatesResponse struct {
	// 总记录数。

	Count *int32 `json:"count,omitempty"`

	Configurations *[]ListConfigurationsResult `json:"configurations,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListConfigurationTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationTemplatesResponse", string(data)}, " ")
}
