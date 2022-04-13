package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckProjectNameRequestV4 struct {
	// 项目名

	ProjectName string `json:"project_name"`
}

func (o CheckProjectNameRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckProjectNameRequestV4 struct{}"
	}

	return strings.Join([]string{"CheckProjectNameRequestV4", string(data)}, " ")
}
