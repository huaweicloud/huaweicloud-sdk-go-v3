package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeDatabaseRequest struct {

	// 实例是否延迟升级，默认false，立即升级。  - true: 延迟升级，实例将在运维时间窗内自动升级。 - false: 立即升级。
	Delay *bool `json:"delay,omitempty"`
}

func (o UpgradeDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabaseRequest struct{}"
	}

	return strings.Join([]string{"UpgradeDatabaseRequest", string(data)}, " ")
}
