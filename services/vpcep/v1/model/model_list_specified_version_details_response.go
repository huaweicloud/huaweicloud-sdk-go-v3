package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecifiedVersionDetailsResponse Response Object
type ListSpecifiedVersionDetailsResponse struct {
	Version        *VersionObject `json:"version,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSpecifiedVersionDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecifiedVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecifiedVersionDetailsResponse", string(data)}, " ")
}
