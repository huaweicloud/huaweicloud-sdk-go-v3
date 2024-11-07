package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeDatabaseVersionResponse Response Object
type BatchUpgradeDatabaseVersionResponse struct {

	// 批量升级结果。
	UpgradeResults *[]UpgradeResult `json:"upgrade_results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchUpgradeDatabaseVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeDatabaseVersionResponse struct{}"
	}

	return strings.Join([]string{"BatchUpgradeDatabaseVersionResponse", string(data)}, " ")
}
