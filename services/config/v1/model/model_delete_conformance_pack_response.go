package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConformancePackResponse Response Object
type DeleteConformancePackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConformancePackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConformancePackResponse struct{}"
	}

	return strings.Join([]string{"DeleteConformancePackResponse", string(data)}, " ")
}
