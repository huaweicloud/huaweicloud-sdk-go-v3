package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelationByIdResultData data，统一的返回结果的最外层数据结构。
type ShowRelationByIdResultData struct {
	Value *RelationVo `json:"value,omitempty"`
}

func (o ShowRelationByIdResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelationByIdResultData struct{}"
	}

	return strings.Join([]string{"ShowRelationByIdResultData", string(data)}, " ")
}
