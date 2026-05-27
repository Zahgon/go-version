// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package version

// Collection is a type that implements the sort.Interface interface
// so that versions can be sorted.
type Collection []*Version

func (v Collection) Len() int { _ = "STUB: not implemented"; return 0 }

func (v Collection) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (v Collection) Swap(i, j int) { _ = "STUB: not implemented"; return }
