package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezeSubCustomersResponse Response Object
type UnfreezeSubCustomersResponse struct {

	// 错误原因，只有部分失败的时候才返回。 具体请参见表1。
	ErrorDetails   *[]CustomerErrorDetail `json:"error_details,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UnfreezeSubCustomersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"UnfreezeSubCustomersResponse", string(data)}, " ")
}
