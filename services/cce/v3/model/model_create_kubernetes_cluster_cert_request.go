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
type CreateKubernetesClusterCertRequest struct {
	ClusterId   string        `json:"cluster_id"`
	ContentType string        `json:"Content-Type"`
	Body        *CertDuration `json:"body,omitempty"`
}

func (o CreateKubernetesClusterCertRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKubernetesClusterCertRequest", string(data)}, " ")
}
