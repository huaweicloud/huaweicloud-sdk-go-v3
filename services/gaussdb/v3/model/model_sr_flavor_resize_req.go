package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SrFlavorResizeReq struct {

	// FE节点CPU、内存规格ID。填空或者不填视为规格ID与原规格ID保持一致。
	FeFlavorId *string `json:"fe_flavor_id,omitempty"`

	// BE节点CPU、内存规格ID。填空或者不填视为规格ID与原规格ID保持一致。
	BeFlavorId *string `json:"be_flavor_id,omitempty"`
}

func (o SrFlavorResizeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrFlavorResizeReq struct{}"
	}

	return strings.Join([]string{"SrFlavorResizeReq", string(data)}, " ")
}
