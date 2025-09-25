package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDataCategoriesRequest Request Object
type UpdateSecurityDataCategoriesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 数据分类id。
	Id string `json:"id"`

	Body *DataCategoryUpdateDto `json:"body,omitempty"`
}

func (o UpdateSecurityDataCategoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataCategoriesRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataCategoriesRequest", string(data)}, " ")
}
