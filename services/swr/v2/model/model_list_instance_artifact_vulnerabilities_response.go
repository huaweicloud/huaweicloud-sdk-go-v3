package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceArtifactVulnerabilitiesResponse Response Object
type ListInstanceArtifactVulnerabilitiesResponse struct {
	Body           map[string]VulnerabilityReports `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListInstanceArtifactVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceArtifactVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceArtifactVulnerabilitiesResponse", string(data)}, " ")
}
