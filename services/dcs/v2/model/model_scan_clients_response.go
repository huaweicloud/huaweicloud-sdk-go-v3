package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanClientsResponse Response Object
type ScanClientsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ScanClientsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanClientsResponse struct{}"
	}

	return strings.Join([]string{"ScanClientsResponse", string(data)}, " ")
}
