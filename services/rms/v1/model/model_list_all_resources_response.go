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
type ListAllResourcesResponse struct {
	// 资源列表
	Resources      *[]ResourceEntity `json:"resources,omitempty"`
	PageInfo       *PageInfo         `json:"page_info,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAllResourcesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAllResourcesResponse", string(data)}, " ")
}
