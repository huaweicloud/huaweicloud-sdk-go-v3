package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyRequest struct {

	// **参数解释**： 授权内容。 - 当授权类型是委托，该字段为委托名称。 - 当授权类型是访问密钥，该字段为访问密钥ID（AK）。  **取值范围**：长度限制64个字符。
	AgencyName string `json:"agency_name"`
}

func (o AgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyRequest struct{}"
	}

	return strings.Join([]string{"AgencyRequest", string(data)}, " ")
}
