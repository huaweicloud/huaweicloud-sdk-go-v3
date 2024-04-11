package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedObjectViewDto struct {

	// 数据实例ID。
	ObjectId *string `json:"objectId,omitempty"`

	// 关联的数据传输对象列表。
	RelatedList *[]BasicObjectQueryViewDto `json:"relatedList,omitempty"`

	// 关联的数据实体列表。
	RelatedEntityList *[]BasicObjectQueryViewDto `json:"relatedEntityList,omitempty"`
}

func (o RelatedObjectViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedObjectViewDto struct{}"
	}

	return strings.Join([]string{"RelatedObjectViewDto", string(data)}, " ")
}
