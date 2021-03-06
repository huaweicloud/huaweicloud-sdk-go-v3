package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type NovaAssociateSecurityGroupRequestBody struct {
	AddSecurityGroup *NovaAddSecurityGroupOption `json:"addSecurityGroup"`
}

func (o NovaAssociateSecurityGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NovaAssociateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NovaAssociateSecurityGroupRequestBody", string(data)}, " ")
}
