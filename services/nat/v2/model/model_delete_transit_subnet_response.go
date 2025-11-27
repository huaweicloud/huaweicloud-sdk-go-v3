package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitSubnetResponse Response Object
type DeleteTransitSubnetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTransitSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitSubnetResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransitSubnetResponse", string(data)}, " ")
}
