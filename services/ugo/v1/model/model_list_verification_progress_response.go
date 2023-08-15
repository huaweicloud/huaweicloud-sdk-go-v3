package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVerificationProgressResponse Response Object
type ListVerificationProgressResponse struct {

	// 对象总数。
	TotalObjectsCount *int32 `json:"total_objects_count,omitempty"`

	// 完成迁移的对象数量。
	CompletedObjectsCount *int32 `json:"completed_objects_count,omitempty"`

	// 对象迁移的对象列表。
	ObjectsList    *[]DatabaseObject `json:"objects_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListVerificationProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVerificationProgressResponse struct{}"
	}

	return strings.Join([]string{"ListVerificationProgressResponse", string(data)}, " ")
}
