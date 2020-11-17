/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 配额信息。
type ShowQuotasRespQuotas struct {
	// 配额列表。
	Resources *[]ShowQuotasRespQuotasResources `json:"resources,omitempty"`
}

func (o ShowQuotasRespQuotas) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotasRespQuotas", string(data)}, " ")
}
