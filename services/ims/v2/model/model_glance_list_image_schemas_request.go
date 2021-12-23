package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceListImageSchemasRequest struct {
}

func (o GlanceListImageSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasRequest", string(data)}, " ")
}
