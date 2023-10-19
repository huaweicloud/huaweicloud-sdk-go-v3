package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyConfigurationsResponse Response Object
type CopyConfigurationsResponse struct {

	// 参数模板ID。
	ConfigurationId *string `json:"configuration_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CopyConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"CopyConfigurationsResponse", string(data)}, " ")
}
