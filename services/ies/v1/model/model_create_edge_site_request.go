package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEdgeSiteRequest struct {
	Body *CreateEdgeSiteRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateEdgeSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeSiteRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeSiteRequest", string(data)}, " ")
}
