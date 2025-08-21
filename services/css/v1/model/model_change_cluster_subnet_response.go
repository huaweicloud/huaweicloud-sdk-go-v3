package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterSubnetResponse Response Object
type ChangeClusterSubnetResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ChangeClusterSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterSubnetResponse struct{}"
	}

	return strings.Join([]string{"ChangeClusterSubnetResponse", string(data)}, " ")
}
