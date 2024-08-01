/*
Copyright 2023 The OpenEBS Authors.

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

package usage

import (
	"strconv"

	"k8s.io/klog/v2"

	ga4Client "github.com/openebs/google-analytics-4/pkg/client"
	ga4Event "github.com/openebs/google-analytics-4/pkg/event"
	k8sapi "github.com/openebs/lib-csi/pkg/client/k8s"
)

// Usage struct represents all information about a usage metric sent to
// Google Analytics with respect to the application
type Usage struct {
	// OpenebsEventBuilder to build the OpenEBSEvent
	OpenebsEventBuilder *ga4Event.OpenebsEventBuilder

	// GA4 Analytics Client
	AnalyticsClient *ga4Client.MeasurementClient
}

// New returns an instance of Usage
func New() *Usage {
	client, err := ga4Client.NewMeasurementClient(
		ga4Client.WithApiSecret(ApiSecret),
		ga4Client.WithMeasurementId(MeasurementId),
	)
	if err != nil {
		return nil
	}
	openebsEventBuilder := ga4Event.NewOpenebsEventBuilder()
	return &Usage{AnalyticsClient: client, OpenebsEventBuilder: openebsEventBuilder}
}

// SetVolumeName i.e pv name
func (u *Usage) SetVolumeName(name string) *Usage {
	u.OpenebsEventBuilder.VolumeName(name)
	return u
}

// SetVolumeClaimName i.e pvc name
func (u *Usage) SetVolumeClaimName(name string) *Usage {
	u.OpenebsEventBuilder.VolumeClaimName(name)
	return u
}

// SetCategory sets the category of an event
func (u *Usage) SetCategory(c string) *Usage {
	u.OpenebsEventBuilder.Category(c)
	return u
}

// SetNodeCount sets the node count for a k8s cluster.
func (u *Usage) SetNodeCount(n string) *Usage {
	u.OpenebsEventBuilder.NodeCount(n)
	return u
}

// SetVolumeCapacity sets the size of a volume.
func (u *Usage) SetVolumeCapacity(volCapG string) *Usage {
	s, _ := toHumanSize(volCapG)
	u.OpenebsEventBuilder.VolumeCapacity(s)
	return u
}

// SetReplicaCount sets the number of replicas for a volume.
func (u *Usage) SetReplicaCount(replicaCount string) *Usage {
	u.OpenebsEventBuilder.ReplicaCount(replicaCount)
	return u
}

// CommonBuild is a common builder method for Usage struct
func (u *Usage) CommonBuild(engineName string) *Usage {
	v := NewVersion()
	_ = v.getVersion(false)

	u.OpenebsEventBuilder.
		Project(AppName).
		EngineInstaller(v.installerType).
		K8sVersion(v.k8sVersion).
		EngineVersion(v.openebsVersion).
		EngineInstaller(v.installerType).
		EngineName(engineName).
		NodeArch(v.nodeArch).
		NodeOs(v.nodeOs).
		NodeKernelVersion(v.nodeKernelVersion)

	return u
}

// ApplicationBuilder Application builder is used for adding k8s&openebs environment detail
// for non install events
func (u *Usage) ApplicationBuilder() *Usage {
	v := NewVersion()
	_ = v.getVersion(false)

	u.AnalyticsClient.SetClientId(v.id)
	u.OpenebsEventBuilder.K8sDefaultNsUid(v.id)

	return u
}

// InstallBuilder is a concrete builder for install events
func (u *Usage) InstallBuilder(override bool) *Usage {
	v := NewVersion()
	clusterSize, _ := k8sapi.NumberOfNodes()
	_ = v.getVersion(override)

	u.AnalyticsClient.SetClientId(v.id)
	u.OpenebsEventBuilder.
		K8sDefaultNsUid(v.id).
		Category(InstallEvent).
		NodeCount(strconv.Itoa(clusterSize))

	return u
}

// Send POSTS an event over to the GA4 API
func (u *Usage) Send() {
	// Instantiate an analytics client
	go func() {
		client := u.AnalyticsClient
		event := u.OpenebsEventBuilder.Build()

		if err := client.Send(event); err != nil {
			klog.Errorf(err.Error())
			return
		}
	}()
}
