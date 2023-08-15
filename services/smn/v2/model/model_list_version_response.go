package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVersionResponse Response Object
type ListVersionResponse struct {
	Version        *VersionItem `json:"version,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionResponse struct{}"
	}

	return strings.Join([]string{"ListVersionResponse", string(data)}, " ")
}
