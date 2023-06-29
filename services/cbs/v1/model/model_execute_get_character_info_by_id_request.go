package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetCharacterInfoByIdRequest Request Object
type ExecuteGetCharacterInfoByIdRequest struct {

	// 形象id
	CharacterId string `json:"character_id"`
}

func (o ExecuteGetCharacterInfoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetCharacterInfoByIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetCharacterInfoByIdRequest", string(data)}, " ")
}
