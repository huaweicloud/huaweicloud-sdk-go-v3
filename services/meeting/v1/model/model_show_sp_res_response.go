package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSpResResponse struct {

	// 已用的企业并发数
	UsedAccountsCount *int32 `json:"usedAccountsCount,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowSpResResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpResResponse struct{}"
	}

	return strings.Join([]string{"ShowSpResResponse", string(data)}, " ")
}
