package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtomicIndexVoDetailData 返回数据。
type AtomicIndexVoDetailData struct {
	Value *AtomicIndexVo `json:"value,omitempty"`
}

func (o AtomicIndexVoDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicIndexVoDetailData struct{}"
	}

	return strings.Join([]string{"AtomicIndexVoDetailData", string(data)}, " ")
}
