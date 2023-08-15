package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBlackWhiteListUsingDeleteResponse Response Object
type DeleteBlackWhiteListUsingDeleteResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteBlackWhiteListUsingDeleteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlackWhiteListUsingDeleteResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlackWhiteListUsingDeleteResponse", string(data)}, " ")
}
