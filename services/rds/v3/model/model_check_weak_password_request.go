package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckWeakPasswordRequest struct {

	// 需要检验的密码。
	Target *string `json:"target,omitempty"`

	// 引擎名称，取值范围：sqlserver, mysql, postgresql，区分大小写。
	Engine *string `json:"engine,omitempty"`
}

func (o CheckWeakPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeakPasswordRequest struct{}"
	}

	return strings.Join([]string{"CheckWeakPasswordRequest", string(data)}, " ")
}
