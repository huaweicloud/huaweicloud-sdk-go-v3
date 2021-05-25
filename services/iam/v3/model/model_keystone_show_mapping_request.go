package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowMappingRequest struct {
	// 待查询的映射ID。

	Id string `json:"id"`
}

func (o KeystoneShowMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowMappingRequest", string(data)}, " ")
}
