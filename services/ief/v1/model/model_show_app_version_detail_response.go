package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAppVersionDetailResponse struct {
	Version        *AppVersionDetail `json:"version,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAppVersionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppVersionDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAppVersionDetailResponse", string(data)}, " ")
}
