package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDisasterDetailResponse struct {

	// 容灾ID
	Id *string `json:"id,omitempty"`

	// 容灾名称
	Name *string `json:"name,omitempty"`

	// 容灾类型
	DrType *string `json:"dr_type,omitempty"`

	// 容灾状态
	Status *string `json:"status,omitempty"`

	PrimaryCluster *DisasterRecoveryCluster `json:"primary_cluster,omitempty"`

	StandbyCluster *DisasterRecoveryCluster `json:"standby_cluster,omitempty"`

	// 容灾同步周期
	DrSyncPeriod *string `json:"dr_sync_period,omitempty"`

	// 容灾启动时间
	StartTime *string `json:"start_time,omitempty"`

	// 容灾创建时间
	CreateTime     *string `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDisasterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDisasterDetailResponse", string(data)}, " ")
}
