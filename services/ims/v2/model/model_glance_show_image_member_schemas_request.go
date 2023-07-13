package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceShowImageMemberSchemasRequest Request Object
type GlanceShowImageMemberSchemasRequest struct {
}

func (o GlanceShowImageMemberSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberSchemasRequest", string(data)}, " ")
}
