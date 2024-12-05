package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareDesignVersionsResultData data，统一的返回结果的最外层数据结构。
type CompareDesignVersionsResultData struct {
	Value *PublishVersionVo `json:"value,omitempty"`
}

func (o CompareDesignVersionsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareDesignVersionsResultData struct{}"
	}

	return strings.Join([]string{"CompareDesignVersionsResultData", string(data)}, " ")
}
