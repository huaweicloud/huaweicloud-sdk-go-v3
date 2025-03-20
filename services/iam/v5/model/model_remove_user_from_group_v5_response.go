package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveUserFromGroupV5Response Response Object
type RemoveUserFromGroupV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveUserFromGroupV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveUserFromGroupV5Response struct{}"
	}

	return strings.Join([]string{"RemoveUserFromGroupV5Response", string(data)}, " ")
}
