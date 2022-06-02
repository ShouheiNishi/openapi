/*
 * Nnef_PFDmanagement Sevice API
 *
 * Packet Flow Description Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_PFDmanagement

import (
	"github.com/free5gc/openapi"
)

// APIClient manages communication with the Nnef_PFDmanagement Sevice API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	IndividualApplicationPFDApi  *IndividualApplicationPFDApiService
	IndividualPFDSubscriptionApi *IndividualPFDSubscriptionApiService
	PFDOfApplicationsApi         *PFDOfApplicationsApiService
	PFDSubscriptionsApi          *PFDSubscriptionsApiService
	NotificationApi              *NotificationApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.httpClient == nil {
		cfg.httpClient = openapi.GetDefaultHttpClient()
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.IndividualApplicationPFDApi = (*IndividualApplicationPFDApiService)(&c.common)
	c.IndividualPFDSubscriptionApi = (*IndividualPFDSubscriptionApiService)(&c.common)
	c.PFDOfApplicationsApi = (*PFDOfApplicationsApiService)(&c.common)
	c.PFDSubscriptionsApi = (*PFDSubscriptionsApiService)(&c.common)
	c.NotificationApi = (*NotificationApiService)(&c.common)

	return c
}
