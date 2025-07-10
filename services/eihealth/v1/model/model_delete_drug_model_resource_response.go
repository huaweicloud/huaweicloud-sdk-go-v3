package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugModelResourceResponse Response Object
type DeleteDrugModelResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugModelResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugModelResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugModelResourceResponse", string(data)}, " ")
}
