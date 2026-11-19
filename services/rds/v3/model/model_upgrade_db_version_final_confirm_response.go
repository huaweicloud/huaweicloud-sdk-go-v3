package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeDbVersionFinalConfirmResponse Response Object
type UpgradeDbVersionFinalConfirmResponse struct {

	// **参数解释**  小版本升级成功后确认完成后生成的异步任务流id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeDbVersionFinalConfirmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDbVersionFinalConfirmResponse struct{}"
	}

	return strings.Join([]string{"UpgradeDbVersionFinalConfirmResponse", string(data)}, " ")
}
