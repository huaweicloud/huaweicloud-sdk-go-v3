package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Exec struct {

	// **参数解释**： 命令行方式采集指标。 **取值范围**： 不涉及。
	Command *[]string `json:"command,omitempty"`
}

func (o Exec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Exec struct{}"
	}

	return strings.Join([]string{"Exec", string(data)}, " ")
}
