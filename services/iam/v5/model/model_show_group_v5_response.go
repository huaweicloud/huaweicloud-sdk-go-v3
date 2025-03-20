package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupV5Response Response Object
type ShowGroupV5Response struct {
	Group          *Group `json:"group,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowGroupV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupV5Response struct{}"
	}

	return strings.Join([]string{"ShowGroupV5Response", string(data)}, " ")
}
