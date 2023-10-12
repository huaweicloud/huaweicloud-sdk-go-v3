package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOneClickAlarmNotificationsResponse Response Object
type UpdateOneClickAlarmNotificationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateOneClickAlarmNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOneClickAlarmNotificationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateOneClickAlarmNotificationsResponse", string(data)}, " ")
}
