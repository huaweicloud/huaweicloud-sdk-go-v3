package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDataobjectRelationsRequestBody 批量关联数据对象请求body体
type BatchCreateDataobjectRelationsRequestBody struct {

	// 关联数据对象的ID列表
	DataobjectIds *[]string `json:"dataobject_ids,omitempty"`

	// 被关联数据对象的ID列表
	RelatedDataobjectIds *[]string `json:"related_dataobject_ids,omitempty"`
}

func (o BatchCreateDataobjectRelationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDataobjectRelationsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateDataobjectRelationsRequestBody", string(data)}, " ")
}
