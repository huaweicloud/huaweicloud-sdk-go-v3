package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddMemberV4Response Response Object
type AddMemberV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AddMemberV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMemberV4Response struct{}"
	}

	return strings.Join([]string{"AddMemberV4Response", string(data)}, " ")
}
