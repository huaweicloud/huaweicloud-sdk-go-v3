package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRomaAppResponse Response Object
type DeleteRomaAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRomaAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRomaAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteRomaAppResponse", string(data)}, " ")
}
