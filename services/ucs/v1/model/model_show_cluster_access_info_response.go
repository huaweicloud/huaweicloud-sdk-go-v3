package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterAccessInfoResponse Response Object
type ShowClusterAccessInfoResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowClusterAccessInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterAccessInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterAccessInfoResponse", string(data)}, " ")
}
