package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobDescription 训练作业描述。
type JobDescription struct {

	// 对训练作业的描述，默认为“NULL”，字符串的长度限制为[0, 256]。
	Description *string `json:"description,omitempty"`
}

func (o JobDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDescription struct{}"
	}

	return strings.Join([]string{"JobDescription", string(data)}, " ")
}
