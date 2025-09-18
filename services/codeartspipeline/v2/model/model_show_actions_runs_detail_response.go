package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowActionsRunsDetailResponse Response Object
type ShowActionsRunsDetailResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowActionsRunsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActionsRunsDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowActionsRunsDetailResponse", string(data)}, " ")
}
