package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Agency struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *AgencyKindObj `json:"kind"`

	Metadata *AgencyMetadata `json:"metadata"`
}

func (o Agency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agency struct{}"
	}

	return strings.Join([]string{"Agency", string(data)}, " ")
}
