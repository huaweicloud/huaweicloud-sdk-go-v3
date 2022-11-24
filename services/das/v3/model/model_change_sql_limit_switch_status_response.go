package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeSqlLimitSwitchStatusResponse struct {

	// SQL限流任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeSqlLimitSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSqlLimitSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeSqlLimitSwitchStatusResponse", string(data)}, " ")
}
