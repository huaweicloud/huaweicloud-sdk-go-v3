package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountInfo struct {

	// **参数解释**： 账号ID，可以通过调用[查询组织账号列表接口]获得，通过返回值中的data.id获得 **约束限制**： 不涉及 **默认取值**： 不涉及
	AccountId string `json:"account_id"`
}

func (o AccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountInfo struct{}"
	}

	return strings.Join([]string{"AccountInfo", string(data)}, " ")
}
