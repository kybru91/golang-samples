// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     [https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample code for sanitizing a model response using model armor along with user prompt.

package modelarmor

// [START modelarmor_sanitize_model_response]

import (
	"context"
	"fmt"
	"io"

	modelarmor "cloud.google.com/go/modelarmor/apiv1"
	modelarmorpb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	"google.golang.org/api/option"
)

// sanitizeModelResponse method sanitizes a model
// response based on the project, location, and template settings.
//
// w io.Writer: The writer to use for logging.
// projectID string: The ID of the project.
// locationID string: The ID of the location.
// templateID string: The ID of the template.
// modelResponse string: The model response to sanitize.
func sanitizeModelResponse(w io.Writer, projectID, locationID, templateID, modelResponse string) error {
	ctx := context.Background()

	// Create options for Model Armor client.
	opts := option.WithEndpoint(fmt.Sprintf("modelarmor.%s.rep.googleapis.com:443", locationID))
	// Create the Model Armor client.
	client, err := modelarmor.NewClient(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed to create client for location %s: %w", locationID, err)
	}
	defer client.Close()

	// Initialize request argument(s)
	modelResponseData := &modelarmorpb.DataItem{
		DataItem: &modelarmorpb.DataItem_Text{
			Text: modelResponse,
		},
	}

	// Prepare request for sanitizing model response.
	req := &modelarmorpb.SanitizeModelResponseRequest{
		Name:              fmt.Sprintf("projects/%s/locations/%s/templates/%s", projectID, locationID, templateID),
		ModelResponseData: modelResponseData,
	}

	// Sanitize the model response.
	response, err := client.SanitizeModelResponse(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to sanitize model response with user prompt for template %s: %w", templateID, err)
	}

	// Sanitization Result.
	fmt.Fprintf(w, "Sanitization Result: %v\n", response)

	return nil
}

// [END modelarmor_sanitize_model_response]
