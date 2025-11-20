package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainConfigurationsResponse Response Object
type ListDomainConfigurationsResponse struct {

	// **参数说明**：域配置列表。
	DomainConfigurations *[]DomainConfigurationDto `json:"domain_configurations,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDomainConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainConfigurationsResponse", string(data)}, " ")
}
