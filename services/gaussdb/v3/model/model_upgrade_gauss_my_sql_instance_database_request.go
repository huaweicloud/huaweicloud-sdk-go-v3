package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpgradeGaussMySqlInstanceDatabaseRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID
	InstanceId string `json:"instance_id"`

	Body *UpgradeDatabaseRequest `json:"body,omitempty"`
}

func (o UpgradeGaussMySqlInstanceDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeGaussMySqlInstanceDatabaseRequest struct{}"
	}

	return strings.Join([]string{"UpgradeGaussMySqlInstanceDatabaseRequest", string(data)}, " ")
}
