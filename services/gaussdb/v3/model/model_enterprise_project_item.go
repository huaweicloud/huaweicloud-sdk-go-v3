package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProjectItem struct {

	// 企业项目ID。
	Id *string `json:"id,omitempty"`

	// 企业项目名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 状态。 - 1：正常。 - 0：异常。
	Status *string `json:"status,omitempty"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Updated *string `json:"updated,omitempty"`
}

func (o EnterpriseProjectItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectItem struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectItem", string(data)}, " ")
}
