// Copyright 2023-2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package bimap

type BiMap[K comparable, V comparable] struct {
	Key2Val map[K]V
	Val2Key map[V]K // reversed
}

func NewBimap[K comparable, V comparable](m map[K]V) BiMap[K, V] {
	reversed := make(map[V]K, len(m))
	for k, v := range m {
		reversed[v] = k
	}
	return BiMap[K, V]{
		Key2Val: m,
		Val2Key: reversed,
	}
}
