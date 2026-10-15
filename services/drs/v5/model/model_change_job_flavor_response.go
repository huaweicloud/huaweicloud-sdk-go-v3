package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeJobFlavorResponse Response Object
type ChangeJobFlavorResponse struct {

	// 规格变更任务ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeJobFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeJobFlavorResponse struct{}"
	}

	return strings.Join([]string{"ChangeJobFlavorResponse", string(data)}, " ")
}
