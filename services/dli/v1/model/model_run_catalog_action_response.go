package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCatalogActionResponse Response Object
type RunCatalogActionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCatalogActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCatalogActionResponse struct{}"
	}

	return strings.Join([]string{"RunCatalogActionResponse", string(data)}, " ")
}
