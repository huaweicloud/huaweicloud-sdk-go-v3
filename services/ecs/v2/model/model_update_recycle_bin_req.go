package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinReq This is a auto create Body Object
type UpdateRecycleBinReq struct {
	RecycleBin *UpdateRecycleBinOption `json:"recycle_bin"`
}

func (o UpdateRecycleBinReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinReq struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinReq", string(data)}, " ")
}
