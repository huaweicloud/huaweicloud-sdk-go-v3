package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateHostAssetValueResponse Response Object
type AssociateHostAssetValueResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateHostAssetValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateHostAssetValueResponse struct{}"
	}

	return strings.Join([]string{"AssociateHostAssetValueResponse", string(data)}, " ")
}
