package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceStsTokenResponseBodyAssumedAgency Contains information about the assumed agency
type GetResourceStsTokenResponseBodyAssumedAgency struct {

	// The URN of the assumed agency
	Urn string `json:"urn"`

	// A unique identifier that contains the agency ID and the agency session name
	Id string `json:"id"`
}

func (o GetResourceStsTokenResponseBodyAssumedAgency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceStsTokenResponseBodyAssumedAgency struct{}"
	}

	return strings.Join([]string{"GetResourceStsTokenResponseBodyAssumedAgency", string(data)}, " ")
}
