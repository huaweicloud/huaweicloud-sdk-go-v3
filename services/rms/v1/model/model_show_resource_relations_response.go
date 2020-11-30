/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowResourceRelationsResponse struct {
	// 资源关系列表
	Relations      *[]ResourceRelation `json:"relations,omitempty"`
	PageInfo       *PageInfo           `json:"page_info,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowResourceRelationsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowResourceRelationsResponse", string(data)}, " ")
}
