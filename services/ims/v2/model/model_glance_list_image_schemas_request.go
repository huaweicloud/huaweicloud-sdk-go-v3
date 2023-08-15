package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceListImageSchemasRequest Request Object
type GlanceListImageSchemasRequest struct {
}

func (o GlanceListImageSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasRequest", string(data)}, " ")
}
