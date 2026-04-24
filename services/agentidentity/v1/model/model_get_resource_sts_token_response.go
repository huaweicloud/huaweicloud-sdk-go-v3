package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceStsTokenResponse Response Object
type GetResourceStsTokenResponse struct {

	// The source identity specified by the principal that is calling the operation
	SourceIdentity *string `json:"source_identity,omitempty"`

	AssumedAgency *GetResourceStsTokenResponseBodyAssumedAgency `json:"assumed_agency,omitempty"`

	Credentials    *GetResourceStsTokenResponseBodyCredentials `json:"credentials,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o GetResourceStsTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceStsTokenResponse struct{}"
	}

	return strings.Join([]string{"GetResourceStsTokenResponse", string(data)}, " ")
}
