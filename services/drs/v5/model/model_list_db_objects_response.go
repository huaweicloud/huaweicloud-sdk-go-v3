package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDbObjectsResponse struct {
	TargetRootDb *TargetRootDb `json:"target_root_db,omitempty"`

	// 数据库对象迁移或同步信息。
	ObjectInfo map[string]DatabaseObject `json:"object_info,omitempty"`

	// 库下表数量的阈值。
	MaxTableNum    *int32 `json:"max_table_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDbObjectsResponse", string(data)}, " ")
}
