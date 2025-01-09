package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOuInfoRequest Request Object
type UpdateOuInfoRequest struct {

	// OU的id。
	OuId string `json:"ou_id"`

	Body *ModifyOuNameInfoV2Req `json:"body,omitempty"`
}

func (o UpdateOuInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOuInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateOuInfoRequest", string(data)}, " ")
}
