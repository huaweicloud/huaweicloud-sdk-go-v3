package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEngineRequest Request Object
type UpgradeEngineRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 升级的引擎Id
	EngineId string `json:"engine_id"`

	Body *EngineUpdateReq `json:"body,omitempty"`
}

func (o UpgradeEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEngineRequest struct{}"
	}

	return strings.Join([]string{"UpgradeEngineRequest", string(data)}, " ")
}
