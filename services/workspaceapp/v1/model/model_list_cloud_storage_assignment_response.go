package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudStorageAssignmentResponse Response Object
type ListCloudStorageAssignmentResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 查询用户存储声明配置。
	Items          *[]CloudStorageAssignment `json:"items,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListCloudStorageAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudStorageAssignmentResponse struct{}"
	}

	return strings.Join([]string{"ListCloudStorageAssignmentResponse", string(data)}, " ")
}
