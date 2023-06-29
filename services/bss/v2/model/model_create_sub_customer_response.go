package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubCustomerResponse Response Object
type CreateSubCustomerResponse struct {

	// 客户ID。 只有成功或者错误码在CBC.99000050时才会返回。
	DomainId *string `json:"domain_id,omitempty"`

	// 客户的华为云账号名。 若请求参数中传递了此参数值，此处返回的响应值与请求参数中取值一致。若请求参数中未传递此参数值，此处返回的响应值为系统随机生成的32位字符串。 只有成功时或者错误码在CBC.99000050时才会返回。
	DomainName     *string `json:"domain_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubCustomerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubCustomerResponse struct{}"
	}

	return strings.Join([]string{"CreateSubCustomerResponse", string(data)}, " ")
}
