package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFeaturesResponse Response Object
type UpdateFeaturesResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFeaturesResponse struct{}"
	}

	return strings.Join([]string{"UpdateFeaturesResponse", string(data)}, " ")
}
