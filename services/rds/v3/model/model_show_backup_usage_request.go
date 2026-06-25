package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupUsageRequest Request Object
type ShowBackupUsageRequest struct {

	// **参数解释**：  引擎名称, 只支持筛选RDS for MySQL, RDS for MariaDB  **约束限制**：  不涉及。  **取值范围**：  - mysql - mariadb  **默认取值**：  不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowBackupUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupUsageRequest", string(data)}, " ")
}
