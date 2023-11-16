package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEngineConfigRequest Request Object
type UpgradeEngineConfigRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 更新的引擎Id
	EngineId string `json:"engine_id"`

	Body *EngineConfigureReq `json:"body,omitempty"`
}

func (o UpgradeEngineConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEngineConfigRequest struct{}"
	}

	return strings.Join([]string{"UpgradeEngineConfigRequest", string(data)}, " ")
}
