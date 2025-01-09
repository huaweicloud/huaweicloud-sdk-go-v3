package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOuInfoResponse Response Object
type UpdateOuInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateOuInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOuInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateOuInfoResponse", string(data)}, " ")
}
