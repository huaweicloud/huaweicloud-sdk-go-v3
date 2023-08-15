package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableComponentResponse Response Object
type EnableComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableComponentResponse struct{}"
	}

	return strings.Join([]string{"EnableComponentResponse", string(data)}, " ")
}
