package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostStatusRequest Request Object
type ShowHostStatusRequest struct {

	// **参数解释：** 域名ID，您可以通过调用查询全部防护域名列表（ListCompositeHosts）获取域名ID。 **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	HostId string `json:"host_id"`
}

func (o ShowHostStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowHostStatusRequest", string(data)}, " ")
}
