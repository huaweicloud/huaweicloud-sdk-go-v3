package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppVersionsResponse struct {
	Version        *AppVersionDetail `json:"version,omitempty" xml:"version"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateAppVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppVersionsResponse struct{}"
	}

	return strings.Join([]string{"CreateAppVersionsResponse", string(data)}, " ")
}
