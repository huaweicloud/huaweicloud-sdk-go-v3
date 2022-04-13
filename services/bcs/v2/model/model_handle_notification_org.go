package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HandleNotificationOrg struct {
	// 加入的组织

	Name string `json:"name"`
}

func (o HandleNotificationOrg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationOrg struct{}"
	}

	return strings.Join([]string{"HandleNotificationOrg", string(data)}, " ")
}
