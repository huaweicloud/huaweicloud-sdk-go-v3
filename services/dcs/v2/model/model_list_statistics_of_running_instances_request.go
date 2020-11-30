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

// Request Object
type ListStatisticsOfRunningInstancesRequest struct {
}

func (o ListStatisticsOfRunningInstancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListStatisticsOfRunningInstancesRequest", string(data)}, " ")
}
