package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugModelResourceResponse Response Object
type CreateDrugModelResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDrugModelResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugModelResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugModelResourceResponse", string(data)}, " ")
}
