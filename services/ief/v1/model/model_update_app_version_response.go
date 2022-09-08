package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAppVersionResponse struct {
	Version        *AppVersionDetail `json:"version,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateAppVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppVersionResponse", string(data)}, " ")
}
