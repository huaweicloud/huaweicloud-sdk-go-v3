package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFsDirQuotaRequestBody 目录配额信息
type CreateFsDirQuotaRequestBody struct {

	// 合法的已存在的目录的全路径
	Path string `json:"path"`

	// 目录的容量大小，单位：MB; 不填写则默认为0，会导致数据无法写入目录; capacity和quota至少二选一。
	Capacity *int32 `json:"capacity,omitempty"`

	// 目录的inode数量限制; 不填写则默认为0，会导致数据无法写入目录; capacity和quota至少二选一。
	Inode *int32 `json:"inode,omitempty"`
}

func (o CreateFsDirQuotaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsDirQuotaRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFsDirQuotaRequestBody", string(data)}, " ")
}
