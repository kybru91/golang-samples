// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secretmanager

// [START secretmanager_delete_secret_label]
import (
	"context"
	"fmt"
	"io"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/genproto/protobuf/field_mask"
)

// deleteSecret updates the metadata about an existing secret and remove an existing label.
func deleteSecretLabel(w io.Writer, name string) error {
	// name := "projects/my-project/secrets/my-secret"

	labelKey := "labelkey"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create secretmanager client: %w", err)
	}
	defer client.Close()

	// Build the request to get the secret.
	req := &secretmanagerpb.GetSecretRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.GetSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	labels := result.Labels

	delete(labels, labelKey)

	// Build the request to update the secret.
	updateReq := &secretmanagerpb.UpdateSecretRequest{
		Secret: &secretmanagerpb.Secret{
			Name:   name,
			Labels: labels,
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"labels"},
		},
	}

	// Call the API.
	updateResult, err := client.UpdateSecret(ctx, updateReq)
	if err != nil {
		return fmt.Errorf("failed to update secret: %w", err)
	}
	fmt.Fprintf(w, "Updated secret: %s\n", updateResult.Name)
	return nil
}

// [END secretmanager_delete_secret_label]
