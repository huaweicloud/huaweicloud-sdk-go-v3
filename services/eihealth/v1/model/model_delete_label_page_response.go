package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLabelPageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLabelPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLabelPageResponse struct{}"
	}

	return strings.Join([]string{"DeleteLabelPageResponse", string(data)}, " ")
}
