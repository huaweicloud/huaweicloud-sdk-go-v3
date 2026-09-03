package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbObjNewResponse Response Object
type DeleteDbObjNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDbObjNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbObjNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbObjNewResponse", string(data)}, " ")
}
