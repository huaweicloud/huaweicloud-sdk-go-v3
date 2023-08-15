package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubnetResponse Response Object
type DeleteSubnetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubnetResponse", string(data)}, " ")
}
