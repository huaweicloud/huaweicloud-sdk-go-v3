package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicZoneLines struct {

	// 线路ID。
	Line *string `json:"line,omitempty"`

	// 线路名称。
	LineName *string `json:"line_name,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o PublicZoneLines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicZoneLines struct{}"
	}

	return strings.Join([]string{"PublicZoneLines", string(data)}, " ")
}
