package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConfigurationsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	Configurations *[]ListConfigurationsResult `json:"configurations,omitempty" xml:"configurations"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}
