// Copyright 2019 Google LLC
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

package subscriptions

// [START pubsub_list_subscriptions]
import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub/v2"
	"cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"google.golang.org/api/iterator"
)

func list(projectID string) ([]*pubsubpb.Subscription, error) {
	// projectID := "my-project-id"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	var subs []*pubsubpb.Subscription
	req := &pubsubpb.ListSubscriptionsRequest{
		Project: fmt.Sprintf("projects/%s", projectID),
	}
	it := client.SubscriptionAdminClient.ListSubscriptions(ctx, req)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Next: %w", err)
		}
		subs = append(subs, s)
	}
	return subs, nil
}

// [END pubsub_list_subscriptions]
