package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportApplicationInstanceServiceProviderMetadataResponse Response Object
type ImportApplicationInstanceServiceProviderMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportApplicationInstanceServiceProviderMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApplicationInstanceServiceProviderMetadataResponse struct{}"
	}

	return strings.Join([]string{"ImportApplicationInstanceServiceProviderMetadataResponse", string(data)}, " ")
}
