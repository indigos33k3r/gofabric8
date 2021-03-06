package fake

import (
	internalversion "github.com/openshift/origin/pkg/oauth/client/clientset_generated/internalclientset/typed/core/internalversion"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCore struct {
	*core.Fake
}

func (c *FakeCore) OAuthClients(namespace string) internalversion.OAuthClientInterface {
	return &FakeOAuthClients{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCore) RESTClient() restclient.Interface {
	var ret *restclient.RESTClient
	return ret
}
