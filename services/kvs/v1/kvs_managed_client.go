package v1

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	sdkConfig "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/common/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"
)

type KvsManagedClient struct {
	KvsClient  IKvsClient
	Endpoint   model.Endpoint
	Credential model.Credential
	IsUsable   atomic.Bool
}

func NewKvsManagedClient(endpoint model.Endpoint, sdkConfig sdkConfig.KvsSdkConfig) (*KvsManagedClient, error) {
	var managedClient KvsManagedClient

	managedClient.Endpoint = endpoint
	managedClient.IsUsable.Store(true)
	managedClient.Credential.Ak = sdkConfig.Credential.Ak
	managedClient.Credential.Sk = sdkConfig.Credential.Sk
	managedClient.Credential.StsToken = sdkConfig.Credential.StsToken

	auth, err := basic.NewCredentialsBuilder().
		WithAk(sdkConfig.Credential.Ak).
		WithSk(sdkConfig.Credential.Sk).
		WithProjectId(sdkConfig.ProjectId).
		SafeBuild()
	if err != nil {
		return nil, err
	}

	httpConfig := config.DefaultHttpConfig().WithTimeout(time.Duration(sdkConfig.ConnectionTimeout) * time.Second)
	if sdkConfig.CaCertificatePath == "" {
		httpConfig.WithIgnoreSSLVerification(true)
	} else {
		caCert, err := ioutil.ReadFile(sdkConfig.CaCertificatePath)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caCert)
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12,
				RootCAs: pool}}
		transport.ResponseHeaderTimeout = time.Duration(sdkConfig.ReadTimeout) * time.Second
		httpConfig.WithHttpTransport(transport)
	}

	endpointUri := "https://" + endpoint.Name + "." + sdkConfig.RegionId + ".myhuaweicloud.com"

	hcClient, err := KvsClientBuilder().
		WithEndpoint(endpointUri).
		WithCredential(auth).
		WithHttpConfig(httpConfig).
		SafeBuild()
	if err != nil {
		return nil, err
	}

	managedClient.KvsClient = NewKvsClient(hcClient)

	return &managedClient, nil
}
