package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteL7PolicyResponse Response Object
type DeleteL7PolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7PolicyResponse", string(data)}, " ")
}
