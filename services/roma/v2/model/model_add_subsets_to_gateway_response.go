package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddSubsetsToGatewayResponse struct {
	Body           *[]Device `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddSubsetsToGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubsetsToGatewayResponse struct{}"
	}

	return strings.Join([]string{"AddSubsetsToGatewayResponse", string(data)}, " ")
}
