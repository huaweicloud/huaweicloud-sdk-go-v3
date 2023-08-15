package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerErrorDetail struct {

	// 返回码。具体请参见状态码。 注意，此时返回的状态码全部为200。
	ErrorCode *string `json:"error_code,omitempty"`

	// 返回码的描述信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 标识ID。该参数对应的是customer_ids。
	Id *string `json:"id,omitempty"`
}

func (o CustomerErrorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerErrorDetail struct{}"
	}

	return strings.Join([]string{"CustomerErrorDetail", string(data)}, " ")
}
