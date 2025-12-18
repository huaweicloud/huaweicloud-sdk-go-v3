package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseVolumeResult struct {

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**: 数据库的缺省表空间名。 **取值范围**: 不涉及。
	TableSpaceName *string `json:"table_space_name,omitempty"`

	// **参数解释**: 表所属用户名称。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 数据库占用空间大小。 **取值范围**: 不涉及。
	DatabaseSize *string `json:"database_size,omitempty"`
}

func (o DatabaseVolumeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseVolumeResult struct{}"
	}

	return strings.Join([]string{"DatabaseVolumeResult", string(data)}, " ")
}
