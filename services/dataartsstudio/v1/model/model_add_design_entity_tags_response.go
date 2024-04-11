package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesignEntityTagsResponse Response Object
type AddDesignEntityTagsResponse struct {
	Data           *TagsResultData `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AddDesignEntityTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesignEntityTagsResponse struct{}"
	}

	return strings.Join([]string{"AddDesignEntityTagsResponse", string(data)}, " ")
}
