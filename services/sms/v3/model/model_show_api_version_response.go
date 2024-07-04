package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiVersionResponse Response Object
type ShowApiVersionResponse struct {

	// API版本号。
	Id *string `json:"id,omitempty"`

	// API链接地址信息。
	Links *[]Link `json:"links,omitempty"`

	// 版本状态。 SUPPORTED表示支持的版本
	Status *string `json:"status,omitempty"`

	// 版本更新时间。 格式为“yyyy-mm-dd Thh:mm:ssZ”。 其中，T指某个时间的开始；Z指UTC时间。例如：2018-09-30T00:00:00Z
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
