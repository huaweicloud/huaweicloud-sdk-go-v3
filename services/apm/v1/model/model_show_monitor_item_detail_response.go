package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorItemDetailResponse Response Object
type ShowMonitorItemDetailResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMonitorItemDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorItemDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorItemDetailResponse", string(data)}, " ")
}
