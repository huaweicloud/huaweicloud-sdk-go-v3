package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecialConfigurationResponse Response Object
type ListSpecialConfigurationResponse struct {

	// 总条数。
	Total *int32 `json:"total,omitempty"`

	// 域名特殊配置信息。
	SpecialConfigurations *[]SpeicialConfiguration `json:"specialConfigurations,omitempty"`
	HttpStatusCode        int                      `json:"-"`
}

func (o ListSpecialConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecialConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ListSpecialConfigurationResponse", string(data)}, " ")
}
