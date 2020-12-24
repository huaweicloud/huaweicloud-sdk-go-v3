/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListStorageTypesResponse struct {
	// 实例磁盘类型信息。
	StorageType *[]Storage `json:"storage_type,omitempty"`
	// 实例专属存储信息。
	DsspoolInfo    *[]DssPoolInfo `json:"dsspool_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStorageTypesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListStorageTypesResponse", string(data)}, " ")
}
