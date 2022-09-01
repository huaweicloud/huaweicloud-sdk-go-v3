package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHostGroupRequest struct {
	Body *DeleteHostGroupRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DeleteHostGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostGroupRequest", string(data)}, " ")
}
