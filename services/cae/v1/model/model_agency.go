package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Agency struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Agency”，该值不可修改。
	Kind string `json:"kind"`

	Metadata *AgencyMetadata `json:"metadata"`
}

func (o Agency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agency struct{}"
	}

	return strings.Join([]string{"Agency", string(data)}, " ")
}
