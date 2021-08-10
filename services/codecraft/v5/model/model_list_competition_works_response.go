package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCompetitionWorksResponse struct {
	// 作品列表

	Works *[]ListWorksResponseModel `json:"works,omitempty"`
	// 作品总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCompetitionWorksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCompetitionWorksResponse struct{}"
	}

	return strings.Join([]string{"ListCompetitionWorksResponse", string(data)}, " ")
}
