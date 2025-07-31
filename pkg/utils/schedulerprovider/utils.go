/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schedulerprovider

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	SupportedSchedulerProviders = sets.New("volcano")
)

const (
	PodGroupNameFmt = "%s-%s-%s"
)

// GetPodGroupName returns the name of the PodGroup for a given LeaderWorkloadSet name, group index and revision (leaderworkerset.sigs.k8s.io/template-revision-hash).
func GetPodGroupName(lwsName, groupIndex, revision string) string {
	return fmt.Sprintf(PodGroupNameFmt, lwsName, groupIndex, revision)
}

// ProviderType defines the type of scheduler provider
type ProviderType string

const (
	Volcano ProviderType = "volcano"
)

// NewSchedulerProvider creates a new scheduler provider based on the type
func NewSchedulerProvider(providerType ProviderType, client client.Client) (SchedulerProvider, error) {
	switch providerType {
	case Volcano:
		return NewVolcanoProvider(client), nil
	default:
		return nil, fmt.Errorf("unsupported scheduler provider type %s, the supported provider list is %v", providerType, SupportedSchedulerProviders.UnsortedList())
	}
}
