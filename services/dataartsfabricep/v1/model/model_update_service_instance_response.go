package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceInstanceResponse Response Object
type UpdateServiceInstanceResponse struct {

	// 实例Id
	Id *string `json:"id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateServiceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateServiceInstanceResponse", string(data)}, " ")
}
