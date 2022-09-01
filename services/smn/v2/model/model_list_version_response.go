package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVersionResponse struct {
	Version        *interface{} `json:"version,omitempty" xml:"version"`
	HttpStatusCode int          `json:"-"`
}

func (o ListVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionResponse struct{}"
	}

	return strings.Join([]string{"ListVersionResponse", string(data)}, " ")
}
