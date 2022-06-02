/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SubscriberDataManagement

import (
	"github.com/free5gc/openapi"
)

// APIClient manages communication with the Nudm_SDM API v2.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AccessAndMobilitySubscriptionDataRetrievalApi   *AccessAndMobilitySubscriptionDataRetrievalApiService
	GPSIToSUPITranslationApi                        *GPSIToSUPITranslationApiService
	ProvidingAcknowledgementOfSteeringOfRoamingApi  *ProvidingAcknowledgementOfSteeringOfRoamingApiService
	ProvidingAcknowledgementOfUEParametersUpdateApi *ProvidingAcknowledgementOfUEParametersUpdateApiService
	RetrievalOfMultipleDataSetsApi                  *RetrievalOfMultipleDataSetsApiService
	RetrievalOfSharedDataApi                        *RetrievalOfSharedDataApiService
	SMFSelectionSubscriptionDataRetrievalApi        *SMFSelectionSubscriptionDataRetrievalApiService
	SMSManagementSubscriptionDataRetrievalApi       *SMSManagementSubscriptionDataRetrievalApiService
	SMSSubscriptionDataRetrievalApi                 *SMSSubscriptionDataRetrievalApiService
	SessionManagementSubscriptionDataRetrievalApi   *SessionManagementSubscriptionDataRetrievalApiService
	SliceSelectionSubscriptionDataRetrievalApi      *SliceSelectionSubscriptionDataRetrievalApiService
	SubscriptionCreationApi                         *SubscriptionCreationApiService
	SubscriptionCreationForSharedDataApi            *SubscriptionCreationForSharedDataApiService
	SubscriptionDeletionApi                         *SubscriptionDeletionApiService
	SubscriptionDeletionForSharedDataApi            *SubscriptionDeletionForSharedDataApiService
	SubscriptionModificationApi                     *SubscriptionModificationApiService
	TraceConfigurationDataRetrievalApi              *TraceConfigurationDataRetrievalApiService
	UEContextInSMFDataRetrievalApi                  *UEContextInSMFDataRetrievalApiService
	UEContextInSMSFDataRetrievalApi                 *UEContextInSMSFDataRetrievalApiService
	DataChangeNotificationCallbackDocumentApi       *DataChangeNotificationCallbackDocumentApiService
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
	c.AccessAndMobilitySubscriptionDataRetrievalApi = (*AccessAndMobilitySubscriptionDataRetrievalApiService)(&c.common)
	c.GPSIToSUPITranslationApi = (*GPSIToSUPITranslationApiService)(&c.common)
	c.ProvidingAcknowledgementOfSteeringOfRoamingApi = (*ProvidingAcknowledgementOfSteeringOfRoamingApiService)(&c.common)
	c.ProvidingAcknowledgementOfUEParametersUpdateApi = (*ProvidingAcknowledgementOfUEParametersUpdateApiService)(&c.common)
	c.RetrievalOfMultipleDataSetsApi = (*RetrievalOfMultipleDataSetsApiService)(&c.common)
	c.RetrievalOfSharedDataApi = (*RetrievalOfSharedDataApiService)(&c.common)
	c.SMFSelectionSubscriptionDataRetrievalApi = (*SMFSelectionSubscriptionDataRetrievalApiService)(&c.common)
	c.SMSManagementSubscriptionDataRetrievalApi = (*SMSManagementSubscriptionDataRetrievalApiService)(&c.common)
	c.SMSSubscriptionDataRetrievalApi = (*SMSSubscriptionDataRetrievalApiService)(&c.common)
	c.SessionManagementSubscriptionDataRetrievalApi = (*SessionManagementSubscriptionDataRetrievalApiService)(&c.common)
	c.SliceSelectionSubscriptionDataRetrievalApi = (*SliceSelectionSubscriptionDataRetrievalApiService)(&c.common)
	c.SubscriptionCreationApi = (*SubscriptionCreationApiService)(&c.common)
	c.SubscriptionCreationForSharedDataApi = (*SubscriptionCreationForSharedDataApiService)(&c.common)
	c.SubscriptionDeletionApi = (*SubscriptionDeletionApiService)(&c.common)
	c.SubscriptionDeletionForSharedDataApi = (*SubscriptionDeletionForSharedDataApiService)(&c.common)
	c.SubscriptionModificationApi = (*SubscriptionModificationApiService)(&c.common)
	c.TraceConfigurationDataRetrievalApi = (*TraceConfigurationDataRetrievalApiService)(&c.common)
	c.UEContextInSMFDataRetrievalApi = (*UEContextInSMFDataRetrievalApiService)(&c.common)
	c.UEContextInSMSFDataRetrievalApi = (*UEContextInSMSFDataRetrievalApiService)(&c.common)
	c.DataChangeNotificationCallbackDocumentApi = (*DataChangeNotificationCallbackDocumentApiService)(&c.common)

	return c
}
