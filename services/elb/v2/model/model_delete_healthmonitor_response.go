package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthmonitorResponse Response Object
type DeleteHealthmonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthmonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorResponse", string(data)}, " ")
}
