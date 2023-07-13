package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpResResponse Response Object
type ShowSpResResponse struct {

	// SP下所有企业已使用的会议并发数量。
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
