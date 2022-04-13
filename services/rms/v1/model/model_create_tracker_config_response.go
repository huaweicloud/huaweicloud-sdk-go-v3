package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTrackerConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTrackerConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateTrackerConfigResponse", string(data)}, " ")
}
