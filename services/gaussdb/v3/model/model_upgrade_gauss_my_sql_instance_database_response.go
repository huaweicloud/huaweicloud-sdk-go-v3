package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeGaussMySqlInstanceDatabaseResponse Response Object
type UpgradeGaussMySqlInstanceDatabaseResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeGaussMySqlInstanceDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeGaussMySqlInstanceDatabaseResponse struct{}"
	}

	return strings.Join([]string{"UpgradeGaussMySqlInstanceDatabaseResponse", string(data)}, " ")
}
