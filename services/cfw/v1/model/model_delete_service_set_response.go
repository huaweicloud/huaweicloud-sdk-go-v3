package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceSetResponse Response Object
type DeleteServiceSetResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteServiceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceSetResponse", string(data)}, " ")
}
