package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComponentConfigurationResponse Response Object
type UpdateComponentConfigurationResponse struct {

	// 结果详情
	Message *string `json:"message,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateComponentConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateComponentConfigurationResponse", string(data)}, " ")
}
