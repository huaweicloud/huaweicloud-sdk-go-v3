package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneratePocketFileResponse Response Object
type GeneratePocketFileResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GeneratePocketFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneratePocketFileResponse struct{}"
	}

	return strings.Join([]string{"GeneratePocketFileResponse", string(data)}, " ")
}
