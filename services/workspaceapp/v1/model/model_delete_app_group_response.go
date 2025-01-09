package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppGroupResponse Response Object
type DeleteAppGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppGroupResponse", string(data)}, " ")
}
