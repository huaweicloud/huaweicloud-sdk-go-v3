/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProjectDemandStaticV4Response struct {
	// 需求统计
	DemandStatistics *[]DemandStatisticResponseV4 `json:"demand_statistics,omitempty"`
}

func (o ListProjectDemandStaticV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectDemandStaticV4Response", string(data)}, " ")
}
