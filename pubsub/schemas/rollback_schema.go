// Copyright 2023 Google LLC
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

package schema

// [START pubsub_rollback_schema]
import (
	"context"
	"fmt"
	"io"

	pubsub "cloud.google.com/go/pubsub/v2/apiv1"
	"cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
)

// rollbackSchema creates a new schema revision that is a copy of the provided revisionID.
func rollbackSchema(w io.Writer, projectID, schemaID, revisionID string) error {
	// projectID := "my-project-id"
	// schemaID := "my-schema"
	// revisionID := "a1b2c3d4"
	ctx := context.Background()
	client, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		return fmt.Errorf("pubsub.NewSchemaClient: %w", err)
	}
	defer client.Close()

	req := &pubsubpb.RollbackSchemaRequest{
		Name:       fmt.Sprintf("projects/%s/schemas/%s", projectID, schemaID),
		RevisionId: revisionID,
	}
	s, err := client.RollbackSchema(ctx, req)
	if err != nil {
		return fmt.Errorf("RollbackSchema: %w", err)
	}
	fmt.Fprintf(w, "Rolled back schema: %#v\n", s)
	return nil
}

// [END pubsub_rollback_schema]
