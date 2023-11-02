package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDatabaseVersionResponse Response Object
type ShowInstanceDatabaseVersionResponse struct {
	Datastore *InstanceDatabaseVersionInfo `json:"datastore,omitempty"`

	// 是否可升级。 - true：是。 - false：否。
	UpgradeFlag    *bool `json:"upgrade_flag,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowInstanceDatabaseVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDatabaseVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDatabaseVersionResponse", string(data)}, " ")
}
