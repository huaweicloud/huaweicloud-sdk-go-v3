package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPreheatingAssetResponse struct {
	// 预热任务数组

	PreheatingResults *[]PreheatingResult `json:"preheating_results,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ShowPreheatingAssetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPreheatingAssetResponse struct{}"
	}

	return strings.Join([]string{"ShowPreheatingAssetResponse", string(data)}, " ")
}
