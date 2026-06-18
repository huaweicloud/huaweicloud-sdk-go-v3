package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSamlProviderV5Response Response Object
type DeleteSamlProviderV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSamlProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSamlProviderV5Response struct{}"
	}

	return strings.Join([]string{"DeleteSamlProviderV5Response", string(data)}, " ")
}
