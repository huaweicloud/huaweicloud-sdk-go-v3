package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReconfigureExtDataSourceActionReq **参数解释**： 更新数据源配置。 **取值范围**： 不涉及。
type ReconfigureExtDataSourceActionReq struct {
	Reconfigure *ReconfigureExtDataSourceAction `json:"reconfigure"`
}

func (o ReconfigureExtDataSourceActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReconfigureExtDataSourceActionReq struct{}"
	}

	return strings.Join([]string{"ReconfigureExtDataSourceActionReq", string(data)}, " ")
}
