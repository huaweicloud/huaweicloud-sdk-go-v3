package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveAccountsRequest Request Object
type BatchRemoveAccountsRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	Body *AccountDto `json:"body,omitempty"`
}

func (o BatchRemoveAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveAccountsRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveAccountsRequest", string(data)}, " ")
}
