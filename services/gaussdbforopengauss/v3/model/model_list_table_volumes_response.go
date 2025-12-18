package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableVolumesResponse Response Object
type ListTableVolumesResponse struct {

	// **参数解释**: 数据库表占用空间列表。
	TableVolumes *[]TableVolumeResult `json:"table_volumes,omitempty"`

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTableVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListTableVolumesResponse", string(data)}, " ")
}
