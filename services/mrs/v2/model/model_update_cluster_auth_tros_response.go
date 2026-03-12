package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterAuthTrosResponse Response Object
type UpdateClusterAuthTrosResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClusterAuthTrosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterAuthTrosResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterAuthTrosResponse", string(data)}, " ")
}
