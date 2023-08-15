package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezeSubCustomersResponse Response Object
type FreezeSubCustomersResponse struct {

	// 错误原因，只有部分失败的时候才返回。 具体请参见表1。
	ErrorDetails   *[]CustomerErrorDetail `json:"error_details,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o FreezeSubCustomersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"FreezeSubCustomersResponse", string(data)}, " ")
}
