package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"os"
	"strings"
)

const (
	xSdkContentSha256  = "X-Sdk-Content-Sha256"
	DefaultIamEndpoint = "https://iam.myhuaweicloud.com"
	IamEndpointEnv     = "HUAWEICLOUD_SDK_IAM_ENDPOINT"
)

type KvsCredentials struct {
	basic.Credentials
	enableBodySignature bool
}

func (s *KvsCredentials) ProcessAuthRequest(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	reqBuilder := req.Builder()
	if !s.enableBodySignature {
		reqBuilder = reqBuilder.AddHeaderParam(xSdkContentSha256, "UNSIGNED-PAYLOAD")
	}
	return s.Credentials.ProcessAuthRequest(client, req)
}

type KvsCredentialsBuilder struct {
	Credentials *KvsCredentials
	errMap      map[string]string
}

func (builder *KvsCredentialsBuilder) WithEnableBodySignature(enableBodySignature bool) *KvsCredentialsBuilder {
	builder.Credentials.enableBodySignature = enableBodySignature
	return builder
}

func NewKvsCredentialsBuilder() *KvsCredentialsBuilder {
	return &KvsCredentialsBuilder{
		Credentials: &KvsCredentials{Credentials: basic.Credentials{IamEndpoint: GetIamEndpoint()}},
		errMap:      make(map[string]string),
	}
}

func GetIamEndpoint() string {
	if endpoint := os.Getenv(IamEndpointEnv); endpoint != "" {
		https := "https://"
		if !strings.HasPrefix(endpoint, https) {
			endpoint = https + endpoint
		}
		return endpoint
	}
	return DefaultIamEndpoint
}

func (builder *KvsCredentialsBuilder) WithIamEndpointOverride(endpoint string) *KvsCredentialsBuilder {
	builder.Credentials.IamEndpoint = endpoint
	return builder
}

func (builder *KvsCredentialsBuilder) WithAk(ak string) *KvsCredentialsBuilder {
	builder.Credentials.AK = ak
	return builder
}

func (builder *KvsCredentialsBuilder) WithSk(sk string) *KvsCredentialsBuilder {
	builder.Credentials.SK = sk
	return builder
}

func (builder *KvsCredentialsBuilder) WithProjectId(projectId string) *KvsCredentialsBuilder {
	builder.Credentials.ProjectId = projectId
	return builder
}

func (builder *KvsCredentialsBuilder) WithSecurityToken(token string) *KvsCredentialsBuilder {
	builder.Credentials.SecurityToken = token
	return builder
}

func (builder *KvsCredentialsBuilder) WithDerivedPredicate(derivedPredicate func(*request.DefaultHttpRequest) bool) *KvsCredentialsBuilder {
	builder.Credentials.DerivedPredicate = derivedPredicate
	return builder
}

func (builder *KvsCredentialsBuilder) WithIdpId(idpId string) *KvsCredentialsBuilder {
	builder.Credentials.IdpId = idpId
	return builder
}

func (builder *KvsCredentialsBuilder) WithIdTokenFile(idTokenFile string) *KvsCredentialsBuilder {
	builder.Credentials.IdTokenFile = idTokenFile
	return builder
}

func (builder *KvsCredentialsBuilder) SafeBuild() (*KvsCredentials, error) {
	if builder.errMap != nil && len(builder.errMap) != 0 {
		errMsg := "build credentials failed: "
		for _, msg := range builder.errMap {
			errMsg += msg + "; "
		}
		return nil, sdkerr.NewCredentialsTypeError(errMsg)
	}

	if builder.Credentials.IdpId != "" || builder.Credentials.IdTokenFile != "" {
		if builder.Credentials.IdpId == "" {
			return nil, sdkerr.NewCredentialsTypeError("IdpId is required when using IdpId&IdTokenFile")
		}
		if builder.Credentials.IdTokenFile == "" {
			return nil, sdkerr.NewCredentialsTypeError("IdTokenFile is required when using IdpId&IdTokenFile")
		}
		if builder.Credentials.ProjectId == "" {
			return nil, sdkerr.NewCredentialsTypeError("ProjectId is required when using IdpId&IdTokenFile")
		}
	}
	return builder.Credentials, nil
}
