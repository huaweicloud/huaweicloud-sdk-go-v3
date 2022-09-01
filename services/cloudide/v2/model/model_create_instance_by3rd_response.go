package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateInstanceBy3rdResponse struct {
	Result *InstancesResponseInstancesVoResult `json:"result,omitempty" xml:"result"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceBy3rdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceBy3rdResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceBy3rdResponse", string(data)}, " ")
}
