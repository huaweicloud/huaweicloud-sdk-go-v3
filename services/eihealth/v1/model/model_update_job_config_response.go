package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobConfigResponse Response Object
type UpdateJobConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateJobConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobConfigResponse", string(data)}, " ")
}
