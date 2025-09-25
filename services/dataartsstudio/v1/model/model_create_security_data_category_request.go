package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDataCategoryRequest Request Object
type CreateSecurityDataCategoryRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DataCategoryInsertDto `json:"body,omitempty"`
}

func (o CreateSecurityDataCategoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityDataCategoryRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityDataCategoryRequest", string(data)}, " ")
}
