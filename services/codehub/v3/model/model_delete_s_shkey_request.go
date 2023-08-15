package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSShkeyRequest Request Object
type DeleteSShkeyRequest struct {

	// sshKey的id
	Id string `json:"id"`
}

func (o DeleteSShkeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSShkeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSShkeyRequest", string(data)}, " ")
}
