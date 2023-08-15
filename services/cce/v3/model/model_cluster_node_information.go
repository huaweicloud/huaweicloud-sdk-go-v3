package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterNodeInformation
type ClusterNodeInformation struct {
	Metadata *ClusterNodeInformationMetadata `json:"metadata"`
}

func (o ClusterNodeInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeInformation struct{}"
	}

	return strings.Join([]string{"ClusterNodeInformation", string(data)}, " ")
}
