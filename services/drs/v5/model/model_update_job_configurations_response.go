package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobConfigurationsResponse Response Object
type UpdateJobConfigurationsResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateJobConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobConfigurationsResponse", string(data)}, " ")
}
