package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateFsDirQuotaResponse struct {

	// 合法的已存在的目录的全路径
	Path *string `json:"path,omitempty"`

	// 目录的容量大小，单位：MB
	Capacity *int32 `json:"capacity,omitempty"`

	// 目录的inode数量限制
	Inode *int32 `json:"inode,omitempty"`

	// 目录已使用的容量大小，单位：MB。仅SFSTurbo HPC型返回该字段
	UsedCapacity *int32 `json:"used_capacity,omitempty"`

	// 目录的已使用的inode数量。仅SFSTurbo HPC型返回该字段
	UsedInode      *int32 `json:"used_inode,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateFsDirQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFsDirQuotaResponse struct{}"
	}

	return strings.Join([]string{"UpdateFsDirQuotaResponse", string(data)}, " ")
}
