package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAgencyRequest Request Object
type CheckAgencyRequest struct {
}

func (o CheckAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAgencyRequest struct{}"
	}

	return strings.Join([]string{"CheckAgencyRequest", string(data)}, " ")
}
