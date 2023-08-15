package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaRemainderData 配额剩余数量信息
type QuotaRemainderData struct {
	Type *QuotaResourceTypeEnum `json:"type,omitempty"`

	// 剩余配额
	Remainder *int32 `json:"remainder,omitempty"`

	// 所需配额
	Need *int32 `json:"need,omitempty"`
}

func (o QuotaRemainderData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaRemainderData struct{}"
	}

	return strings.Join([]string{"QuotaRemainderData", string(data)}, " ")
}
