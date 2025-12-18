package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmDetailResponse Response Object
type ShowDdmDetailResponse struct {

	// 实例详情。
	Instances      *[]InstanceDetail `json:"instances,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowDdmDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDdmDetailResponse", string(data)}, " ")
}
