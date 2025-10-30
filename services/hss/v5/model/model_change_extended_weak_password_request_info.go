package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeExtendedWeakPasswordRequestInfo 修改镜像的自定义弱口令
type ChangeExtendedWeakPasswordRequestInfo struct {

	// **参数解释**: 自定义弱口令，选填，可编辑 **取值范围**: 最小值0，最大值300
	ExtendedWeakPassword *[]string `json:"extended_weak_password,omitempty"`
}

func (o ChangeExtendedWeakPasswordRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeExtendedWeakPasswordRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeExtendedWeakPasswordRequestInfo", string(data)}, " ")
}
