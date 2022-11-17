package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeProtectEipResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ChangeProtectEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeProtectEipResponse struct{}"
	}

	return strings.Join([]string{"ChangeProtectEipResponse", string(data)}, " ")
}
