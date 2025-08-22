package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ReportCollection represents the ReportCollection schema from the OpenAPI specification
type ReportCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []ReportRecordContract `json:"value,omitempty"` // Page values.
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
}

// ReportRecordContract represents the ReportRecordContract schema from the OpenAPI specification
type ReportRecordContract struct {
	Region string `json:"region,omitempty"` // Country region to which this record data is related.
	Cachehitcount int `json:"cacheHitCount,omitempty"` // Number of times when content was served from cache policy.
	Callcountblocked int `json:"callCountBlocked,omitempty"` // Number of calls blocked due to invalid credentials. This includes calls returning HttpStatusCode.Unauthorized and HttpStatusCode.Forbidden and HttpStatusCode.TooManyRequests
	Interval string `json:"interval,omitempty"` // Length of aggregation period. Interval must be multiple of 15 minutes and may not be zero. The value should be in ISO 8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).
	Timestamp string `json:"timestamp,omitempty"` // Start of aggregation period. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Apitimemin float64 `json:"apiTimeMin,omitempty"` // Minimum time it took to process request.
	Callcountfailed int `json:"callCountFailed,omitempty"` // Number of calls failed due to proxy or backend errors. This includes calls returning HttpStatusCode.BadRequest(400) and any Code between HttpStatusCode.InternalServerError (500) and 600
	Name string `json:"name,omitempty"` // Name depending on report endpoint specifies product, API, operation or developer name.
	Subscriptionid string `json:"subscriptionId,omitempty"` // Subscription identifier path. /subscriptions/{subscriptionId}
	Apiregion string `json:"apiRegion,omitempty"` // API region identifier.
	Apitimeavg float64 `json:"apiTimeAvg,omitempty"` // Average time it took to process request.
	Servicetimemax float64 `json:"serviceTimeMax,omitempty"` // Maximum time it took to process request on backend.
	Apiid string `json:"apiId,omitempty"` // API identifier path. /apis/{apiId}
	Zip string `json:"zip,omitempty"` // Zip code to which this record data is related.
	Callcountsuccess int `json:"callCountSuccess,omitempty"` // Number of successful calls. This includes calls returning HttpStatusCode <= 301 and HttpStatusCode.NotModified and HttpStatusCode.TemporaryRedirect
	Callcounttotal int `json:"callCountTotal,omitempty"` // Total number of calls.
	Productid string `json:"productId,omitempty"` // Product identifier path. /products/{productId}
	Servicetimeavg float64 `json:"serviceTimeAvg,omitempty"` // Average time it took to process request on backend.
	Apitimemax float64 `json:"apiTimeMax,omitempty"` // Maximum time it took to process request.
	Cachemisscount int `json:"cacheMissCount,omitempty"` // Number of times content was fetched from backend.
	Country string `json:"country,omitempty"` // Country to which this record data is related.
	Servicetimemin float64 `json:"serviceTimeMin,omitempty"` // Minimum time it took to process request on backend.
	Callcountother int `json:"callCountOther,omitempty"` // Number of other calls.
	Userid string `json:"userId,omitempty"` // User identifier path. /users/{userId}
	Bandwidth int64 `json:"bandwidth,omitempty"` // Bandwidth consumed.
	Operationid string `json:"operationId,omitempty"` // Operation identifier path. /apis/{apiId}/operations/{operationId}
}

// RequestReportCollection represents the RequestReportCollection schema from the OpenAPI specification
type RequestReportCollection struct {
	Value []RequestReportRecordContract `json:"value,omitempty"` // Page values.
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
}

// RequestReportRecordContract represents the RequestReportRecordContract schema from the OpenAPI specification
type RequestReportRecordContract struct {
	Ipaddress string `json:"ipAddress,omitempty"` // The client IP address associated with this request.
	Operationid string `json:"operationId,omitempty"` // Operation identifier path. /apis/{apiId}/operations/{operationId}
	Apitime float64 `json:"apiTime,omitempty"` // The total time it took to process this request.
	Userid string `json:"userId,omitempty"` // User identifier path. /users/{userId}
	Requestsize int `json:"requestSize,omitempty"` // The size of this request..
	Responsecode int `json:"responseCode,omitempty"` // The HTTP status code returned by the gateway.
	Servicetime float64 `json:"serviceTime,omitempty"` // he time it took to forward this request to the backend and get the response back.
	Url string `json:"url,omitempty"` // The full URL associated with this request.
	Responsesize int `json:"responseSize,omitempty"` // The size of the response returned by the gateway.
	Apiregion string `json:"apiRegion,omitempty"` // Azure region where the gateway that processed this request is located.
	Backendresponsecode string `json:"backendResponseCode,omitempty"` // The HTTP status code received by the gateway as a result of forwarding this request to the backend.
	Cache string `json:"cache,omitempty"` // Specifies if response cache was involved in generating the response. If the value is none, the cache was not used. If the value is hit, cached response was returned. If the value is miss, the cache was used but lookup resulted in a miss and request was fulfilled by the backend.
	Productid string `json:"productId,omitempty"` // Product identifier path. /products/{productId}
	Subscriptionid string `json:"subscriptionId,omitempty"` // Subscription identifier path. /subscriptions/{subscriptionId}
	Timestamp string `json:"timestamp,omitempty"` // The date and time when this request was received by the gateway in ISO 8601 format.
	Apiid string `json:"apiId,omitempty"` // API identifier path. /apis/{apiId}
	Method string `json:"method,omitempty"` // The HTTP method associated with this request..
	Requestid string `json:"requestId,omitempty"` // Request Identifier.
}
