package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCatalogActionRequest Request Object
type RunCatalogActionRequest struct {
	Body *RunCatalogActionRequestBody `json:"body,omitempty"`
}

func (o RunCatalogActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCatalogActionRequest struct{}"
	}

	return strings.Join([]string{"RunCatalogActionRequest", string(data)}, " ")
}
