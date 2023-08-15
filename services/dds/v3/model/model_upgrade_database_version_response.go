package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeDatabaseVersionResponse Response Object
type UpgradeDatabaseVersionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeDatabaseVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabaseVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeDatabaseVersionResponse", string(data)}, " ")
}
