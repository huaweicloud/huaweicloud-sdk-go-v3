package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyInstanceConfigurationsResponse Response Object
type CopyInstanceConfigurationsResponse struct {

	// 参数模板ID。
	ConfigurationId *string `json:"configuration_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CopyInstanceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"CopyInstanceConfigurationsResponse", string(data)}, " ")
}
