package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStorageTypesResponse struct {

	// 实例磁盘类型信息。
	StorageType    *[]Storage `json:"storage_type,omitempty" xml:"storage_type"`
	HttpStatusCode int        `json:"-"`
}

func (o ListStorageTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypesResponse struct{}"
	}

	return strings.Join([]string{"ListStorageTypesResponse", string(data)}, " ")
}
