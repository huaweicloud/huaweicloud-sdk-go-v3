package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSpecifiedVersionResponse struct {
	Version        *Versions `json:"version,omitempty" xml:"version"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowSpecifiedVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecifiedVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowSpecifiedVersionResponse", string(data)}, " ")
}
