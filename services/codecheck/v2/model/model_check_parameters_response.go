package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckParametersResponse struct {
	Body           *[]CheckParametersRes `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CheckParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckParametersResponse struct{}"
	}

	return strings.Join([]string{"CheckParametersResponse", string(data)}, " ")
}
