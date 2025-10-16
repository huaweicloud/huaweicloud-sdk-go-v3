package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeDatabaseVersionResponse Response Object
type UpgradeDatabaseVersionResponse struct {

	// **参数解释：** 任务ID。 **取值范围：** 不涉及。
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
