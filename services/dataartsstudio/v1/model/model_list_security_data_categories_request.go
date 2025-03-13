package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDataCategoriesRequest Request Object
type ListSecurityDataCategoriesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSecurityDataCategoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDataCategoriesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDataCategoriesRequest", string(data)}, " ")
}
