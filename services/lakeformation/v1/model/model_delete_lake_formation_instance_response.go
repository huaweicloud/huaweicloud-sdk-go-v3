package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLakeFormationInstanceResponse Response Object
type DeleteLakeFormationInstanceResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLakeFormationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLakeFormationInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteLakeFormationInstanceResponse", string(data)}, " ")
}
