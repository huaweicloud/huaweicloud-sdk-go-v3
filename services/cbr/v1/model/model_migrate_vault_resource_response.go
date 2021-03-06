package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type MigrateVaultResourceResponse struct {
	//

	MigratedResources *[]string `json:"migrated_resources,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o MigrateVaultResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceResponse", string(data)}, " ")
}
