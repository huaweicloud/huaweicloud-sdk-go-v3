package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFsDirQuotaRequestBody 目录配额信息
type UpdateFsDirQuotaRequestBody struct {

	// 合法的已存在的目录的全路径
	Path string `json:"path"`

	// 目录的容量大小，单位：MB
	Capacity *int32 `json:"capacity,omitempty"`

	// 目录的inode数量限制
	Inode *int32 `json:"inode,omitempty"`
}

func (o UpdateFsDirQuotaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFsDirQuotaRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFsDirQuotaRequestBody", string(data)}, " ")
}
