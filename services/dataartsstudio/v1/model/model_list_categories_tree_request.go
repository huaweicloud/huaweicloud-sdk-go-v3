package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCategoriesTreeRequest Request Object
type ListCategoriesTreeRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	WorkspaceId string `json:"workspace_id"`
}

func (o ListCategoriesTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCategoriesTreeRequest struct{}"
	}

	return strings.Join([]string{"ListCategoriesTreeRequest", string(data)}, " ")
}
