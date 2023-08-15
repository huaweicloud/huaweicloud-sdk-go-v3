package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateMappingResponse Response Object
type KeystoneUpdateMappingResponse struct {
	Mapping        *MappingResult `json:"mapping,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneUpdateMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateMappingResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateMappingResponse", string(data)}, " ")
}
