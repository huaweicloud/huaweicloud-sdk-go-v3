package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dis/v2/model"
)

type DisClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDisClient(hcClient *http_client.HcHttpClient) *DisClient {
	return &DisClient{HcClient: hcClient}
}

func DisClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}
