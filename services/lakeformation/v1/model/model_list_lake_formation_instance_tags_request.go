package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLakeFormationInstanceTagsRequest Request Object
type ListLakeFormationInstanceTagsRequest struct {

	// 使用预定义标签，true表示使用
	UsePredefineTags bool `json:"use_predefine_tags"`
}

func (o ListLakeFormationInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLakeFormationInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLakeFormationInstanceTagsRequest", string(data)}, " ")
}
