package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddDevDataResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDevDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDevDataResponse struct{}"
	}

	return strings.Join([]string{"AddDevDataResponse", string(data)}, " ")
}
