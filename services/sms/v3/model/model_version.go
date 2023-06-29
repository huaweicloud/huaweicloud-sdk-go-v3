package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Version API版本
type Version struct {

	// API版本号。
	Id *string `json:"id,omitempty"`

	// API链接地址信息。
	Links *[]Link `json:"links,omitempty"`

	// 版本状态。 取值“CURRENT”，表示该版本为主推版本。 取值“DEPRECATED”，表示为废弃版本，存在后续删除的可能。
	Status *string `json:"status,omitempty"`

	// 版本更新时间。 格式为“yyyy-mm-dd Thh:mm:ssZ”。 其中，T指某个时间的开始；Z指UTC时间。例如：2018-09-30T00:00:00Z
	Updated *string `json:"updated,omitempty"`
}

func (o Version) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Version struct{}"
	}

	return strings.Join([]string{"Version", string(data)}, " ")
}
