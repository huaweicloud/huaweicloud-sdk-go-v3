package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstanceDefaultResponse Response Object
type UpdateLakeFormationInstanceDefaultResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLakeFormationInstanceDefaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceDefaultResponse struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceDefaultResponse", string(data)}, " ")
}
