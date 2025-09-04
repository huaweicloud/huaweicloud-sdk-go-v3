package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudStorageAssignmentCreateInfo 创建文件夹返回信息。
type CloudStorageAssignmentCreateInfo struct {

	// 存储分配的唯一标识符。
	Id *string `json:"id,omitempty"`

	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 域ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 文件夹名称。
	FolderName *string `json:"folder_name,omitempty"`

	// 用户名称。
	Attach *string `json:"attach,omitempty"`

	// 用户ID。
	AttachId *string `json:"attach_id,omitempty"`

	AttachType *AttachType `json:"attach_type,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 是否创建成功。
	IsSuccess *bool `json:"is_success,omitempty"`
}

func (o CloudStorageAssignmentCreateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudStorageAssignmentCreateInfo struct{}"
	}

	return strings.Join([]string{"CloudStorageAssignmentCreateInfo", string(data)}, " ")
}
