package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageAssignmentResponse Response Object
type ListStorageAssignmentResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 查询用户存储声明配置。
	Items          *[]PersistentStorageAssignment `json:"items,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListStorageAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageAssignmentResponse struct{}"
	}

	return strings.Join([]string{"ListStorageAssignmentResponse", string(data)}, " ")
}
