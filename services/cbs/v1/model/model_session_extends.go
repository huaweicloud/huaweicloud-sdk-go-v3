package model

import (
	"encoding/json"

	"strings"
)

//
type SessionExtends struct {
	TagIds *Tag `json:"tag_ids,omitempty"`
	// 领域列表，多个领域用分号隔开。如果设置了领域且领域不为空，就从这些领域中匹配答案，否则就从该用户的全部知识库匹配答案。  当前最多支持10个领域。

	DomainIds *[]string `json:"domain_ids,omitempty"`
	// 问题来源 其他支持用户自定义，最终体现在问答日志里

	Source *string `json:"source,omitempty"`
}

func (o SessionExtends) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SessionExtends struct{}"
	}

	return strings.Join([]string{"SessionExtends", string(data)}, " ")
}
