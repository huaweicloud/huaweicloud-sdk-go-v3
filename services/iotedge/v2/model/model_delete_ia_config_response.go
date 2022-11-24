package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteIaConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIaConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIaConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteIaConfigResponse", string(data)}, " ")
}
