package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubjectRequest Request Object
type CreateSubjectRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *CatalogParamsVo `json:"body,omitempty"`
}

func (o CreateSubjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubjectRequest struct{}"
	}

	return strings.Join([]string{"CreateSubjectRequest", string(data)}, " ")
}
