package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIndicatorsResponse Response Object
type ImportIndicatorsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ImportIndicatorsResponse", string(data)}, " ")
}
