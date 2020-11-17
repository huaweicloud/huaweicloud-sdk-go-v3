/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type UpdateSinkTaskQuotaReq struct {
	// 转储任务的总个数。
	SinkMaxTasks string `json:"sink_max_tasks"`
}

func (o UpdateSinkTaskQuotaReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSinkTaskQuotaReq", string(data)}, " ")
}
