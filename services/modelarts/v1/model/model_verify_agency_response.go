package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyAgencyResponse Response Object
type VerifyAgencyResponse struct {

	// 是否有权使用该委托，true表示有权，false表示无权
	IsAuthorized *bool `json:"is_authorized,omitempty"`

	// **参数解释**： 授权内容。 - 当授权类型是委托，该字段为委托名称。 - 当授权类型是访问密钥，该字段为访问密钥ID（AK）。  **取值范围**：长度限制64个字符。
	AgencyName *string `json:"agency_name,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o VerifyAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyAgencyResponse struct{}"
	}

	return strings.Join([]string{"VerifyAgencyResponse", string(data)}, " ")
}
