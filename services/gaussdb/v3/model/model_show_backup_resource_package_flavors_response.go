package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupResourcePackageFlavorsResponse Response Object
type ShowBackupResourcePackageFlavorsResponse struct {

	// **参数解释**：  可用备份资源包规格列表。
	Flavors        *[]BackupResourcePackageFlavor `json:"flavors,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowBackupResourcePackageFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupResourcePackageFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupResourcePackageFlavorsResponse", string(data)}, " ")
}
