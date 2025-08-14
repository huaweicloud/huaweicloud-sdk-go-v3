package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudStorageAssignmentRequest Request Object
type ListCloudStorageAssignmentRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	// 存储声明的类型,目前只支持USER,后面可以拓展。 * `USER` -  个人文件夹 * `SHARE` -  共享文件夹
	ClaimMode string `json:"claim_mode"`

	// 被分配的个体或组的名称，当传attach_names时，本字段不生效。
	Attach *string `json:"attach,omitempty"`

	// 被分配的个体或组的名称。
	AttachNames *[]string `json:"attach_names,omitempty"`

	// 是否查询容量过滤： - true : 是。 - false: 否。
	CapacityFilter *string `json:"capacity_filter,omitempty"`
}

func (o ListCloudStorageAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudStorageAssignmentRequest struct{}"
	}

	return strings.Join([]string{"ListCloudStorageAssignmentRequest", string(data)}, " ")
}
