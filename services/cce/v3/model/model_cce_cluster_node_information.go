/*
 * cce
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type CceClusterNodeInformation struct {
	Metadata *CceClusterNodeInformationMetadata `json:"metadata"`
}

func (o CceClusterNodeInformation) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CceClusterNodeInformation", string(data)}, " ")
}
