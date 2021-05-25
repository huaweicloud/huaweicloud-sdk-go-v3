package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListStacksByTagRequest struct {
	// 技术栈标签。默认为空值，查询全部。 目前可取值 Java,Maven,CPP,Vue,ServiceComb,Blockchain,GO, Node.js,DCN,Quantum,JavaScript,Ruby,Python;可查询多个标签

	Tags *[]string `json:"tags,omitempty"`
}

func (o ListStacksByTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListStacksByTagRequest struct{}"
	}

	return strings.Join([]string{"ListStacksByTagRequest", string(data)}, " ")
}
