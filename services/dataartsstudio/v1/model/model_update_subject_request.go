package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubjectRequest Request Object
type UpdateSubjectRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *CatalogParamsVo `json:"body,omitempty"`
}

func (o UpdateSubjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubjectRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubjectRequest", string(data)}, " ")
}
