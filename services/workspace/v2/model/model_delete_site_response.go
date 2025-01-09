package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSiteResponse Response Object
type DeleteSiteResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSiteResponse struct{}"
	}

	return strings.Join([]string{"DeleteSiteResponse", string(data)}, " ")
}
