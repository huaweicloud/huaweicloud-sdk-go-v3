package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicZoneLines struct {

	// **参数解释：** 线路ID。 **取值范围：** 不涉及。
	Line *string `json:"line,omitempty"`

	// **参数解释：** 线路名称。 **取值范围：** 不涉及。
	LineName *string `json:"line_name,omitempty"`

	// **参数解释：** 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o PublicZoneLines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicZoneLines struct{}"
	}

	return strings.Join([]string{"PublicZoneLines", string(data)}, " ")
}
