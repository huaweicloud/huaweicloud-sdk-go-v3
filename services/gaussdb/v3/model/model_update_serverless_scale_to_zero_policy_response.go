package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessScaleToZeroPolicyResponse Response Object
type UpdateServerlessScaleToZeroPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerlessScaleToZeroPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessScaleToZeroPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerlessScaleToZeroPolicyResponse", string(data)}, " ")
}
