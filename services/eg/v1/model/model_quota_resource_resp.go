package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaResourceResp struct {

	// 错误描述
	Resources *[]QuotaItemInfo `json:"resources,omitempty"`
}

func (o QuotaResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourceResp struct{}"
	}

	return strings.Join([]string{"QuotaResourceResp", string(data)}, " ")
}
