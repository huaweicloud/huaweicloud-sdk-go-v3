package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppVersionsResponse Response Object
type CreateAppVersionsResponse struct {
	Version        *AppVersionDetail `json:"version,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateAppVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppVersionsResponse struct{}"
	}

	return strings.Join([]string{"CreateAppVersionsResponse", string(data)}, " ")
}
