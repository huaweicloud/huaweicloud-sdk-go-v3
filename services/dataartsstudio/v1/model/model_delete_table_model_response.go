package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableModelResponse Response Object
type DeleteTableModelResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteTableModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableModelResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableModelResponse", string(data)}, " ")
}
