package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDesignEntityTagsResponse Response Object
type RemoveDesignEntityTagsResponse struct {
	Data           *TagsResultData `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RemoveDesignEntityTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDesignEntityTagsResponse struct{}"
	}

	return strings.Join([]string{"RemoveDesignEntityTagsResponse", string(data)}, " ")
}
