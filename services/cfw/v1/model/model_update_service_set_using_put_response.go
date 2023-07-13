package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceSetUsingPutResponse Response Object
type UpdateServiceSetUsingPutResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateServiceSetUsingPutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSetUsingPutResponse struct{}"
	}

	return strings.Join([]string{"UpdateServiceSetUsingPutResponse", string(data)}, " ")
}
