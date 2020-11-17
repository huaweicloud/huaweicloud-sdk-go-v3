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

// Request Object
type UpdateSinkTaskQuotaRequest struct {
	ProjectId   string                  `json:"project_id"`
	ConnectorId string                  `json:"connector_id"`
	Body        *UpdateSinkTaskQuotaReq `json:"body,omitempty"`
}

func (o UpdateSinkTaskQuotaRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSinkTaskQuotaRequest", string(data)}, " ")
}
