package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaDetail 当前用户云堡垒机的配额信息。返回默认值null。
type QuotaDetail struct {

	// 中文配额描述。
	ZhCn string `json:"zh_cn"`

	// 英文配额描述。
	EnUs string `json:"en_us"`

	// 租户剩余配额数量。
	Remaining int32 `json:"remaining"`
}

func (o QuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
