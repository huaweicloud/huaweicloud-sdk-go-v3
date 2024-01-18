package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageAssignmentRequest Request Object
type ListStorageAssignmentRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	// 存储声明的类型，claim_mode为share时，storage_claim_id必填。 * `USER` -  用户目录 * `SHARE` - 共享目录
	ClaimMode string `json:"claim_mode"`

	// WKS存储目录声明ID。
	StorageClaimId *string `json:"storage_claim_id,omitempty"`

	// 成员。
	Attach *string `json:"attach,omitempty"`

	// 关联对象类型： * `USER` -  用户 * `USER_GROUP` - 用户组
	AttachType *string `json:"attach_type,omitempty"`
}

func (o ListStorageAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageAssignmentRequest struct{}"
	}

	return strings.Join([]string{"ListStorageAssignmentRequest", string(data)}, " ")
}
