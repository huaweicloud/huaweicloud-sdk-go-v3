package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SharePersistentStorageClaim 基于WKS存储创建的，文件夹存储声明
type SharePersistentStorageClaim struct {

	// WKS存储目录声明ID
	StorageClaimId *string `json:"storage_claim_id,omitempty"`

	// 存储对象路径 注: path是对象在系统中的完整路径 例如系统中存在如下目录结构的数据. SFS-Tmp: └─shares   ├─image   └─video image的路径: shares/image/ video的路径: shares/video/
	FolderPath *string `json:"folder_path,omitempty"`

	// 路径分隔符
	Delimiter *string `json:"delimiter,omitempty"`

	ClaimMode *ClaimMode `json:"claim_mode,omitempty"`

	// 共享目录的关联用户、用户组数量
	Count map[string]int32 `json:"count,omitempty"`
}

func (o SharePersistentStorageClaim) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SharePersistentStorageClaim struct{}"
	}

	return strings.Join([]string{"SharePersistentStorageClaim", string(data)}, " ")
}
