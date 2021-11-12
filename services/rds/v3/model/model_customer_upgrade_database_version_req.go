package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerUpgradeDatabaseVersionReq struct {
	// 是否延迟升级。

	Delay *bool `json:"delay,omitempty"`
}

func (o CustomerUpgradeDatabaseVersionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerUpgradeDatabaseVersionReq struct{}"
	}

	return strings.Join([]string{"CustomerUpgradeDatabaseVersionReq", string(data)}, " ")
}
