package webhook

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/date"
    "github.com/Azure/go-autorest/autorest/to"
    "net/http"
)

// CreateOrUpdateParameters is the parameters supplied to the create or update
// webhook operation.
type CreateOrUpdateParameters struct {
    Name *string `json:"name,omitempty"`
    *CreateOrUpdateProperties `json:"properties,omitempty"`
}

// CreateOrUpdateProperties is the properties of the create webhook operation.
type CreateOrUpdateProperties struct {
    IsEnabled *bool `json:"isEnabled,omitempty"`
    URI *string `json:"uri,omitempty"`
    ExpiryTime *date.Time `json:"expiryTime,omitempty"`
    Parameters *map[string]*string `json:"parameters,omitempty"`
    Runbook *RunbookAssociationProperty `json:"runbook,omitempty"`
    RunOn *string `json:"runOn,omitempty"`
}

// ErrorResponse is error response of an operation failure
type ErrorResponse struct {
    Code *string `json:"code,omitempty"`
    Message *string `json:"message,omitempty"`
}

// ListResult is the response model for the list webhook operation.
type ListResult struct {
    autorest.Response `json:"-"`
    Value *[]Model `json:"value,omitempty"`
    NextLink *string `json:"nextLink,omitempty"`
}

// ListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListResult) ListResultPreparer() (*http.Request, error) {
    if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
        return nil, nil
    }
    return autorest.Prepare(&http.Request{},
        autorest.AsJSON(),
        autorest.AsGet(),
        autorest.WithBaseURL(to.String(client.NextLink)));
}

// Model is definition of the webhook type.
type Model struct {
    autorest.Response `json:"-"`
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    *Properties `json:"properties,omitempty"`
}

// Properties is definition of the webhook properties
type Properties struct {
    IsEnabled *bool `json:"isEnabled,omitempty"`
    URI *string `json:"uri,omitempty"`
    ExpiryTime *date.Time `json:"expiryTime,omitempty"`
    LastInvokedTime *date.Time `json:"lastInvokedTime,omitempty"`
    Parameters *map[string]*string `json:"parameters,omitempty"`
    Runbook *RunbookAssociationProperty `json:"runbook,omitempty"`
    RunOn *string `json:"runOn,omitempty"`
    CreationTime *date.Time `json:"creationTime,omitempty"`
    LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
    Description *string `json:"description,omitempty"`
}

// RunbookAssociationProperty is the runbook property associated with the
// entity.
type RunbookAssociationProperty struct {
    Name *string `json:"name,omitempty"`
}

// String is
type String struct {
    autorest.Response `json:"-"`
    Value *string `json:"value,omitempty"`
}

// UpdateParameters is the parameters supplied to the update webhook operation.
type UpdateParameters struct {
    Name *string `json:"name,omitempty"`
    *UpdateProperties `json:"properties,omitempty"`
}

// UpdateProperties is the properties of the update webhook.
type UpdateProperties struct {
    IsEnabled *bool `json:"isEnabled,omitempty"`
    RunOn *string `json:"runOn,omitempty"`
    Parameters *map[string]*string `json:"parameters,omitempty"`
    Description *string `json:"description,omitempty"`
}
