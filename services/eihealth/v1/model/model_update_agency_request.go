package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyRequest Request Object
type UpdateAgencyRequest struct {
}

func (o UpdateAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyRequest", string(data)}, " ")
}
