// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by main. DO NOT EDIT.

package loadbalancer

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *L3n4Addr) deepEqual(other *L3n4Addr) bool {
	if other == nil {
		return false
	}

	if in.L4Addr != other.L4Addr {
		return false
	}

	if in.Scope != other.Scope {
		return false
	}

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *L3n4AddrID) deepEqual(other *L3n4AddrID) bool {
	if other == nil {
		return false
	}

	if !in.L3n4Addr.DeepEqual(&other.L3n4Addr) {
		return false
	}

	if in.ID != other.ID {
		return false
	}

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *L4Addr) deepEqual(other *L4Addr) bool {
	if other == nil {
		return false
	}

	if in.Protocol != other.Protocol {
		return false
	}
	if in.Port != other.Port {
		return false
	}

	return true
}
