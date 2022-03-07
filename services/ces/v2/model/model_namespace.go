package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Namespace struct {
	// 指标的命名空间指标命名空间，例如弹性云服务器命名空间(格式为service.item；service和item必须是字符串，以字母开头，可包含0-9/a-z/A-Z/_；长度范围[3,32]。)

	Namespace *string `json:"namespace,omitempty"`
}

func (o Namespace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Namespace struct{}"
	}

	return strings.Join([]string{"Namespace", string(data)}, " ")
}
