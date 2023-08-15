package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveLakeFormationInstanceOutRecycleBinResponse Response Object
type MoveLakeFormationInstanceOutRecycleBinResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MoveLakeFormationInstanceOutRecycleBinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveLakeFormationInstanceOutRecycleBinResponse struct{}"
	}

	return strings.Join([]string{"MoveLakeFormationInstanceOutRecycleBinResponse", string(data)}, " ")
}
