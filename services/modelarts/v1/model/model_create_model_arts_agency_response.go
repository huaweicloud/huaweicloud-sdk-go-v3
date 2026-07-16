package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelArtsAgencyResponse Response Object
type CreateModelArtsAgencyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateModelArtsAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelArtsAgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateModelArtsAgencyResponse", string(data)}, " ")
}
