package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMcpServerResponse Response Object
type UpdateMcpServerResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMcpServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMcpServerResponse struct{}"
	}

	return strings.Join([]string{"UpdateMcpServerResponse", string(data)}, " ")
}
