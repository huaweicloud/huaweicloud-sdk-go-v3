/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAddonTemplatesRequest struct {
	ContentType       string  `json:"Content-Type"`
	AddonTemplateName *string `json:"addon_template_name,omitempty"`
}

func (o ListAddonTemplatesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAddonTemplatesRequest", string(data)}, " ")
}
