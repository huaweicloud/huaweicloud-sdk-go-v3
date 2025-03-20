package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddUserToGroupV5Response Response Object
type AddUserToGroupV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AddUserToGroupV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserToGroupV5Response struct{}"
	}

	return strings.Join([]string{"AddUserToGroupV5Response", string(data)}, " ")
}
