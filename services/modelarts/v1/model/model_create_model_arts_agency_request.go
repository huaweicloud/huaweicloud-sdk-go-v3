package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelArtsAgencyRequest Request Object
type CreateModelArtsAgencyRequest struct {
	Body *ModelArtsAgencyRequest `json:"body,omitempty"`
}

func (o CreateModelArtsAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelArtsAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateModelArtsAgencyRequest", string(data)}, " ")
}
