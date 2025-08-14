package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateProfileResponse Response Object
type DisassociateProfileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateProfileResponse struct{}"
	}

	return strings.Join([]string{"DisassociateProfileResponse", string(data)}, " ")
}
