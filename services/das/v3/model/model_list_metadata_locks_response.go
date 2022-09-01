package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMetadataLocksResponse struct {

	// 元数据锁列表
	MetadataLocks *[]MetadataLock `json:"metadata_locks,omitempty" xml:"metadata_locks"`

	// 元数据锁数量
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMetadataLocksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadataLocksResponse struct{}"
	}

	return strings.Join([]string{"ListMetadataLocksResponse", string(data)}, " ")
}
