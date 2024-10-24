package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlternateIdentifierDto A unique identifier for a user or group that is not the its primary identifier.This value can be an identifier from an external identity provider (IdP) that is associated with the group or a unique attribute.
type AlternateIdentifierDto struct {
	ExternalId *ExternalIdDto `json:"external_id,omitempty"`

	UniqueAttribute *UniqueAttributeDto `json:"unique_attribute,omitempty"`
}

func (o AlternateIdentifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlternateIdentifierDto struct{}"
	}

	return strings.Join([]string{"AlternateIdentifierDto", string(data)}, " ")
}
