package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReconfigureExtDataSourceActionReq 更新数据源配置
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
