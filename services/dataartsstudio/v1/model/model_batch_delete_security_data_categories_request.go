package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityDataCategoriesRequest Request Object
type BatchDeleteSecurityDataCategoriesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchDeleteCategoryGroupDto `json:"body,omitempty"`
}

func (o BatchDeleteSecurityDataCategoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityDataCategoriesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityDataCategoriesRequest", string(data)}, " ")
}
