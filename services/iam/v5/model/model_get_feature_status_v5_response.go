package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetFeatureStatusV5Response Response Object
type GetFeatureStatusV5Response struct {
	FeatureStatus  *FeatureStatus `json:"feature_status,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o GetFeatureStatusV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFeatureStatusV5Response struct{}"
	}

	return strings.Join([]string{"GetFeatureStatusV5Response", string(data)}, " ")
}
