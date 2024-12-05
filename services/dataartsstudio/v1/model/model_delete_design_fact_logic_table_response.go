package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignFactLogicTableResponse Response Object
type DeleteDesignFactLogicTableResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignFactLogicTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignFactLogicTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignFactLogicTableResponse", string(data)}, " ")
}
