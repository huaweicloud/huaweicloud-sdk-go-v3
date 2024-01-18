package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PersistentStorageAssignment 存储声明关联作用对象。
type PersistentStorageAssignment struct {

	// WKS存储目录声明ID。
	StorageClaimId *string `json:"storage_claim_id,omitempty"`

	// 存储对象路径。 注: path是对象在系统中的完整路径。 例如系统中存在如下目录结构的数据。 SFS-Tmp: └─shares   ├─image   └─video image的路径: shares/image/ video的路径: shares/video/
	FolderPath *string `json:"folder_path,omitempty"`

	// 路径分隔符。
	Delimiter *string `json:"delimiter,omitempty"`

	ClaimMode *ClaimMode `json:"claim_mode,omitempty"`

	StorageMetadata *StorageMetadata `json:"storage_metadata,omitempty"`

	PolicyStatement *PolicyStatement `json:"policy_statement,omitempty"`

	Attachment *Attachment `json:"attachment,omitempty"`
}

func (o PersistentStorageAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentStorageAssignment struct{}"
	}

	return strings.Join([]string{"PersistentStorageAssignment", string(data)}, " ")
}
