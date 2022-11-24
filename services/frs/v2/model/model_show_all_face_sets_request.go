package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAllFaceSetsRequest struct {
}

func (o ShowAllFaceSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllFaceSetsRequest struct{}"
	}

	return strings.Join([]string{"ShowAllFaceSetsRequest", string(data)}, " ")
}
