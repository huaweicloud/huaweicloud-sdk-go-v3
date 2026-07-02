package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLockWaitSessionResponse Response Object
type ShowLockWaitSessionResponse struct {
	AbnormalRootCause *AbnormalRootCause `json:"abnormal_root_cause,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ShowLockWaitSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockWaitSessionResponse struct{}"
	}

	return strings.Join([]string{"ShowLockWaitSessionResponse", string(data)}, " ")
}
