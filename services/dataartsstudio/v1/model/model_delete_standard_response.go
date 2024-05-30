package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStandardResponse Response Object
type DeleteStandardResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteStandardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStandardResponse struct{}"
	}

	return strings.Join([]string{"DeleteStandardResponse", string(data)}, " ")
}
