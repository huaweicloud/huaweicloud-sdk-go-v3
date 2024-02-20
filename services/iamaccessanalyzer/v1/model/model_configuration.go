package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Configuration struct {
	IamAgency *IamAgency `json:"iam_agency,omitempty"`

	ObsBucket *ObsBucket `json:"obs_bucket,omitempty"`
}

func (o Configuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configuration struct{}"
	}

	return strings.Join([]string{"Configuration", string(data)}, " ")
}
