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

type ClusterInformation struct {
	Spec *ClusterInformationSpec `json:"spec"`
}

func (o ClusterInformation) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ClusterInformation", string(data)}, " ")
}
