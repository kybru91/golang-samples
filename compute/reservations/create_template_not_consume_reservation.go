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

package snippets

// [START compute_template_not_consume_reservation]
import (
	"context"
	"fmt"
	"io"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
)

// createInstanceNotConsumeReservation creates a new instance template, which won't consume reservations
func createTemplateNotConsumeReservation(w io.Writer, projectID, templateName string) error {
	// projectID := "your_project_id"
	// templateName := "your_template_name"

	ctx := context.Background()
	instanceTemplatesClient, err := compute.NewInstanceTemplatesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstanceTemplatesRESTClient: %w", err)
	}
	defer instanceTemplatesClient.Close()

	req := &computepb.InsertInstanceTemplateRequest{
		Project: projectID,
		InstanceTemplateResource: &computepb.InstanceTemplate{
			Name: proto.String(templateName),
			Properties: &computepb.InstanceProperties{
				// The template describes the size and source image of the boot disk
				// to attach to the instance.
				Disks: []*computepb.AttachedDisk{
					{
						InitializeParams: &computepb.AttachedDiskInitializeParams{
							DiskSizeGb:  proto.Int64(250),
							SourceImage: proto.String("projects/debian-cloud/global/images/family/debian-11"),
						},
						AutoDelete: proto.Bool(true),
						Boot:       proto.Bool(true),
					},
				},
				MachineType: proto.String("e2-standard-4"),
				// The template connects the instance to the `default` network,
				// without specifying a subnetwork.
				NetworkInterfaces: []*computepb.NetworkInterface{
					{
						Name: proto.String("global/networks/default"),
						// The template lets the instance use an external IP address.
						AccessConfigs: []*computepb.AccessConfig{
							{
								Name:        proto.String("External NAT"),
								Type:        proto.String(computepb.AccessConfig_ONE_TO_ONE_NAT.String()),
								NetworkTier: proto.String(computepb.AccessConfig_PREMIUM.String()),
							},
						},
					},
				},
				ReservationAffinity: &computepb.ReservationAffinity{
					ConsumeReservationType: proto.String("NO_RESERVATION"),
				},
			},
		},
	}

	op, err := instanceTemplatesClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to create instance template: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance template created\n")

	return nil
}

// [END compute_template_not_consume_reservation]
