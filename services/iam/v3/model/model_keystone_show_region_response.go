package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowRegionResponse Response Object
type KeystoneShowRegionResponse struct {
	Region         *Region `json:"region,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneShowRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowRegionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowRegionResponse", string(data)}, " ")
}
