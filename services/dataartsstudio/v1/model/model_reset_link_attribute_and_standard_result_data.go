package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetLinkAttributeAndStandardResultData data，统一的返回结果的最外层数据结构。
type ResetLinkAttributeAndStandardResultData struct {
	Value *LinkAttributeAndElementVo `json:"value,omitempty"`
}

func (o ResetLinkAttributeAndStandardResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLinkAttributeAndStandardResultData struct{}"
	}

	return strings.Join([]string{"ResetLinkAttributeAndStandardResultData", string(data)}, " ")
}
