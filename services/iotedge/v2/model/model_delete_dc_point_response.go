package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcPointResponse Response Object
type DeleteDcPointResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDcPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcPointResponse struct{}"
	}

	return strings.Join([]string{"DeleteDcPointResponse", string(data)}, " ")
}
