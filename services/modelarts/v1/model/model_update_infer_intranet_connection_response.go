package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInferIntranetConnectionResponse Response Object
type UpdateInferIntranetConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInferIntranetConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInferIntranetConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateInferIntranetConnectionResponse", string(data)}, " ")
}
