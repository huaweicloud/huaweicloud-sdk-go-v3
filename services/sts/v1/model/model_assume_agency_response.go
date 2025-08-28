package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyResponse Response Object
type AssumeAgencyResponse struct {

	// 调用链里最初调用者所声明的身份。
	SourceIdentity *string `json:"source_identity,omitempty"`

	AssumedAgency *AssumedAgencyDto `json:"assumed_agency,omitempty"`

	Credentials    *CredentialsDto `json:"credentials,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AssumeAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyResponse struct{}"
	}

	return strings.Join([]string{"AssumeAgencyResponse", string(data)}, " ")
}
