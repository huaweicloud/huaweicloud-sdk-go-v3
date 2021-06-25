package model

import (
	"encoding/json"

	"strings"
)

type UpdateCoverByThumbnailReq struct {
	// 截图url<br/>

	ThumbnailUrl string `json:"thumbnailUrl"`
}

func (o UpdateCoverByThumbnailReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCoverByThumbnailReq struct{}"
	}

	return strings.Join([]string{"UpdateCoverByThumbnailReq", string(data)}, " ")
}
