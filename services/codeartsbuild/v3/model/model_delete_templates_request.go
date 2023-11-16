package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplatesRequest Request Object
type DeleteTemplatesRequest struct {

	// uuid
	Uuid string `json:"uuid"`
}

func (o DeleteTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplatesRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesRequest", string(data)}, " ")
}
