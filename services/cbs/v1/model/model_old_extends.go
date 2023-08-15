package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OldExtends
type OldExtends struct {

	// 领域列表，多个领域用分号隔开。如果设置了领域且领域不为空，就从这些领域中匹配答案，否则就从该用户的全部知识库匹配答案。  当前最多支持10个领域。
	Domains *[]string `json:"domains,omitempty"`

	// 返回答案数量，默认为5，取值范围1~10。
	Top *int32 `json:"top,omitempty"`
}

func (o OldExtends) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OldExtends struct{}"
	}

	return strings.Join([]string{"OldExtends", string(data)}, " ")
}
