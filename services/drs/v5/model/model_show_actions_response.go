package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowActionsResponse Response Object
type ShowActionsResponse struct {
	JobAction      *JobActions `json:"job_action,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActionsResponse struct{}"
	}

	return strings.Join([]string{"ShowActionsResponse", string(data)}, " ")
}
