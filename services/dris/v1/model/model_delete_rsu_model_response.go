package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRsuModelResponse Response Object
type DeleteRsuModelResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRsuModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRsuModelResponse struct{}"
	}

	return strings.Join([]string{"DeleteRsuModelResponse", string(data)}, " ")
}
