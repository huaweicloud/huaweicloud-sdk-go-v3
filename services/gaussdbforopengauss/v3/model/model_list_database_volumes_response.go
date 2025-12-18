package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseVolumesResponse Response Object
type ListDatabaseVolumesResponse struct {

	// **参数解释**: 数据库占用空间列表。
	DatabaseVolumes *[]DatabaseVolumeResult `json:"database_volumes,omitempty"`

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseVolumesResponse", string(data)}, " ")
}
