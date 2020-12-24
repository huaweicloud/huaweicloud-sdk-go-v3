/*
 * AS
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListScalingTagInfosByTenantIdResponse struct {
	// 资源标签。
	Tags           *[]TagsMultiValue `json:"tags,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListScalingTagInfosByTenantIdResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListScalingTagInfosByTenantIdResponse", string(data)}, " ")
}
