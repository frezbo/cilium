// SPDX-License-Identifier: Apache-2.0
// Copyright 2019-2020 Authors of Cilium

package k8s

import (
	"path"
	"time"

	"github.com/cilium/cilium/pkg/kvstore"

	"k8s.io/client-go/tools/cache"
)

// CCNPStatusesPath is the KVStore key prefix for CCNP status
var CCNPStatusesPath = path.Join(kvstore.BaseKeyPrefix, "state", "ccnpstatuses", "v2")

// CCNPStatusEventHandler handles status updates events for all the CCNPs
// in the cluster. Upon creation of Clusterwide policies, it will start a
// controller for that CNP which handles sending of updates for that CCNP to
// the kubernetes API server. Upon receiving eventes from the key-value store
// it will send the update for the CCNP corresponding to the status update to
// the controller for that CCNP.
type CCNPStatusEventHandler struct {
	*CNPStatusEventHandler
}

// NewCCNPStatusEventHandler returns a new CCNPStatusEventHandler.
// which is more or less a wrapper around the CNPStatusEventHandler itself.
func NewCCNPStatusEventHandler(k8sStore cache.Store, updateInterval time.Duration) *CCNPStatusEventHandler {
	return &CCNPStatusEventHandler{
		CNPStatusEventHandler: NewCNPStatusEventHandler(k8sStore, updateInterval),
	}
}
