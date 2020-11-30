/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListStatisticsOfRunningInstancesResponse struct {
	// 该租户下处于“运行中”状态的实例的统计信息。
	Statistics     *[]InstanceStatistic `json:"statistics,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListStatisticsOfRunningInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListStatisticsOfRunningInstancesResponse", string(data)}, " ")
}
