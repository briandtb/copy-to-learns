package route

import (
	"context"
	"hotrod/pkg/log"
	"hotrod/pkg/tracing"
	"net/http"
	"net/url"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Client is a remote client that implements route.Interface
type Client struct {
	tracer   opentracing.Tracer
	logger   log.Factory
	client   *tracing.HTTPClient
	hostPort string
}

// NewClient creates a new route.Client
func NewClient(tracer opentracing.Tracer, logger log.Factory, hostPort string) *Client {
	return &Client{
		tracer: tracer,
		logger: logger,
		client: &tracing.HTTPClient{
			Client: &http.Client{Transport: &nethttp.Transport{}},
		},
		hostPort: hostPort,
	}
}

// FindRoute implements route.Interface#FindRoute as an RPC
func (c *Client) FindRoute(ctx context.Context, pickup, dropoff string) (*Route, error) {
	c.logger.For(ctx).Info("Finding route", zap.String("pickup", pickup), zap.String("dropoff", dropoff))

	v := url.Values{}
	v.Set("pickup", pickup)
	v.Set("dropoff", dropoff)
	url := "http://" + c.hostPort + "/route?" + v.Encode()
	var route Route
	if err := c.client.GetJSON(ctx, "/route", url, &route); err != nil {
		c.logger.For(ctx).Error("Error getting route", zap.Error(err))
		return nil, err
	}
	return &route, nil
}
