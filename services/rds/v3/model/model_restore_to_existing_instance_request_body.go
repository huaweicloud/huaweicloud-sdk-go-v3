/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type RestoreToExistingInstanceRequestBody struct {
	Source *RestoreToExistingInstanceRequestBodySource `json:"source"`
	Target *RestoreToExistingInstanceRequestBodyTarget `json:"target"`
}

func (o RestoreToExistingInstanceRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreToExistingInstanceRequestBody", string(data)}, " ")
}
