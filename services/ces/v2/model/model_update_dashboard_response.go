package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDashboardResponse Response Object
type UpdateDashboardResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDashboardResponse struct{}"
	}

	return strings.Join([]string{"UpdateDashboardResponse", string(data)}, " ")
}
