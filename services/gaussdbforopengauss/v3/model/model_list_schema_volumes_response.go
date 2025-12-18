package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemaVolumesResponse Response Object
type ListSchemaVolumesResponse struct {

	// **参数解释**: 数据库schema占用空间列表
	SchemaVolumes *[]SchemaVolumeResult `json:"schema_volumes,omitempty"`

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSchemaVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListSchemaVolumesResponse", string(data)}, " ")
}
