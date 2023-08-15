package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowIdentityProviderRequest Request Object
type KeystoneShowIdentityProviderRequest struct {

	// 待查询的身份提供商ID。
	Id string `json:"id"`
}

func (o KeystoneShowIdentityProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowIdentityProviderRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowIdentityProviderRequest", string(data)}, " ")
}
