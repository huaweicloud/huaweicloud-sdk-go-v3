package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitSubnetTagResponse Response Object
type DeleteTransitSubnetTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTransitSubnetTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitSubnetTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransitSubnetTagResponse", string(data)}, " ")
}
