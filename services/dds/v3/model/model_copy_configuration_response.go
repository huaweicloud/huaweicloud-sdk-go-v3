package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyConfigurationResponse Response Object
type CopyConfigurationResponse struct {

	// 复制后的参数模板ID。
	ConfigurationId *string `json:"configuration_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CopyConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CopyConfigurationResponse", string(data)}, " ")
}
