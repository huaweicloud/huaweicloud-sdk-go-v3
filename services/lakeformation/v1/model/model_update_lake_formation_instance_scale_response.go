package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstanceScaleResponse Response Object
type UpdateLakeFormationInstanceScaleResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLakeFormationInstanceScaleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceScaleResponse struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceScaleResponse", string(data)}, " ")
}
