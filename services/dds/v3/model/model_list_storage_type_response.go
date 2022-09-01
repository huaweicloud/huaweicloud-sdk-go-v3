package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStorageTypeResponse struct {

	// 实例磁盘类型信息。
	StorageType *[]Storage `json:"storage_type,omitempty" xml:"storage_type"`

	// 实例专属存储信息。
	DssPoolInfo    *[]DssPoolInfo `json:"dss_pool_info,omitempty" xml:"dss_pool_info"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStorageTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypeResponse struct{}"
	}

	return strings.Join([]string{"ListStorageTypeResponse", string(data)}, " ")
}
