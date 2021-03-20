package model

import (
	"encoding/json"

	"strings"
)

type ListServerGroupsPageInfoResult struct {
	//

	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ListServerGroupsPageInfoResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServerGroupsPageInfoResult struct{}"
	}

	return strings.Join([]string{"ListServerGroupsPageInfoResult", string(data)}, " ")
}
