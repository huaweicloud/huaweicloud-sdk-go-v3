package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostSatisfactionReq struct {
	// 满意度评分，当前仅支持二级评分，1表示满意，-1表示不满意。

	Degree int32 `json:"degree"`
}

func (o PostSatisfactionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostSatisfactionReq struct{}"
	}

	return strings.Join([]string{"PostSatisfactionReq", string(data)}, " ")
}
