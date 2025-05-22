package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinPolicyResponse Response Object
type UpdateRecycleBinPolicyResponse struct {
	RecycleBin     *RecycleBinResponseBody `json:"recycle_bin,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateRecycleBinPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinPolicyResponse", string(data)}, " ")
}
