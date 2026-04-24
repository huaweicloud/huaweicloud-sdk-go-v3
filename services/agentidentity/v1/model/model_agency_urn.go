package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyUrn The URN of the agency used to obtain IAM temporary credentials.
type AgencyUrn struct {
}

func (o AgencyUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyUrn struct{}"
	}

	return strings.Join([]string{"AgencyUrn", string(data)}, " ")
}
