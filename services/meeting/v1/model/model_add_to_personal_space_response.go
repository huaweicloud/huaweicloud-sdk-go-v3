package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddToPersonalSpaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddToPersonalSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddToPersonalSpaceResponse struct{}"
	}

	return strings.Join([]string{"AddToPersonalSpaceResponse", string(data)}, " ")
}
