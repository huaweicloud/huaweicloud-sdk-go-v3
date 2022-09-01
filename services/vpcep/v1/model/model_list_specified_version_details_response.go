package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSpecifiedVersionDetailsResponse struct {
	Version        *Version `json:"version,omitempty" xml:"version"`
	HttpStatusCode int      `json:"-"`
}

func (o ListSpecifiedVersionDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecifiedVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecifiedVersionDetailsResponse", string(data)}, " ")
}
