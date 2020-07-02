// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
)

const _ = "template_collectProbeOuter"

func collectProbeOuter_false(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.headID[batchSize-1]
	for i := hj.probeState.prevBatchResumeIdx; i < batchSize; i++ {
		currentID := hj.ht.probeScratch.headID[i]

		for {
			if nResults >= hj.outputBatchSize {
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
				return nResults
			}

			hj.probeState.probeRowUnmatched[nResults] = currentID == 0
			if currentID > 0 {
				hj.probeState.buildIdx[nResults] = int(currentID - 1)
			} else {
				// If currentID == 0, then probeRowUnmatched will have been set - and
				// we set the corresponding buildIdx to zero so that (as long as the
				// build hash table has at least one row) we can copy the values vector
				// without paying attention to probeRowUnmatched.
				hj.probeState.buildIdx[nResults] = 0
			}
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			currentID = hj.ht.same[currentID]
			hj.ht.probeScratch.headID[i] = currentID
			nResults++

			if currentID == 0 {
				break
			}
		}
	}
	return nResults
}

func collectProbeOuter_true(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.headID[batchSize-1]
	_ = sel[batchSize-1]
	for i := hj.probeState.prevBatchResumeIdx; i < batchSize; i++ {
		currentID := hj.ht.probeScratch.headID[i]

		for {
			if nResults >= hj.outputBatchSize {
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
				return nResults
			}

			hj.probeState.probeRowUnmatched[nResults] = currentID == 0
			if currentID > 0 {
				hj.probeState.buildIdx[nResults] = int(currentID - 1)
			} else {
				// If currentID == 0, then probeRowUnmatched will have been set - and
				// we set the corresponding buildIdx to zero so that (as long as the
				// build hash table has at least one row) we can copy the values vector
				// without paying attention to probeRowUnmatched.
				hj.probeState.buildIdx[nResults] = 0
			}
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			currentID = hj.ht.same[currentID]
			hj.ht.probeScratch.headID[i] = currentID
			nResults++

			if currentID == 0 {
				break
			}
		}
	}
	return nResults
}

const _ = "template_collectProbeNoOuter"

func collectProbeNoOuter_false(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.headID[batchSize-1]
	for i := hj.probeState.prevBatchResumeIdx; i < batchSize; i++ {
		currentID := hj.ht.probeScratch.headID[i]
		for currentID != 0 {
			if nResults >= hj.outputBatchSize {
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
				return nResults
			}

			hj.probeState.buildIdx[nResults] = int(currentID - 1)
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			currentID = hj.ht.same[currentID]
			hj.ht.probeScratch.headID[i] = currentID
			nResults++
		}
	}
	return nResults
}

func collectProbeNoOuter_true(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.headID[batchSize-1]
	_ = sel[batchSize-1]
	for i := hj.probeState.prevBatchResumeIdx; i < batchSize; i++ {
		currentID := hj.ht.probeScratch.headID[i]
		for currentID != 0 {
			if nResults >= hj.outputBatchSize {
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
				return nResults
			}

			hj.probeState.buildIdx[nResults] = int(currentID - 1)
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			currentID = hj.ht.same[currentID]
			hj.ht.probeScratch.headID[i] = currentID
			nResults++
		}
	}
	return nResults
}

// This code snippet collects the "matches" for LEFT ANTI and EXCEPT ALL joins.
// "Matches" are in quotes because we're actually interested in non-matches
// from the left side.
const _ = "template_collectAnti"

func collectAnti_false(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.headID[batchSize-1]
	for i := int(0); i < batchSize; i++ {
		currentID := hj.ht.probeScratch.headID[i]
		if currentID == 0 {
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				// currentID of 0 indicates that ith probing row didn't have a match, so
				// we include it into the output.
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

func collectAnti_true(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.headID[batchSize-1]
	_ = sel[batchSize-1]
	for i := int(0); i < batchSize; i++ {
		currentID := hj.ht.probeScratch.headID[i]
		if currentID == 0 {
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				// currentID of 0 indicates that ith probing row didn't have a match, so
				// we include it into the output.
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

const _ = "template_distinctCollectProbeOuter"

func distinctCollectProbeOuter_false(hj *hashJoiner, batchSize int, sel []int) {
	// Early bounds checks.
	_ = hj.ht.probeScratch.groupID[batchSize-1]
	_ = hj.probeState.probeRowUnmatched[batchSize-1]
	_ = hj.probeState.buildIdx[batchSize-1]
	_ = hj.probeState.probeIdx[batchSize-1]
	for i := int(0); i < batchSize; i++ {
		// Index of keys and outputs in the hash table is calculated as ID - 1.
		id := hj.ht.probeScratch.groupID[i]
		rowUnmatched := id == 0
		hj.probeState.probeRowUnmatched[i] = rowUnmatched
		if !rowUnmatched {
			hj.probeState.buildIdx[i] = int(id - 1)
		}
		{
			var __retval_0 int
			{
				{
					__retval_0 = i
				}
			}
			hj.probeState.probeIdx[i] = __retval_0
		}
	}
}

func distinctCollectProbeOuter_true(hj *hashJoiner, batchSize int, sel []int) {
	// Early bounds checks.
	_ = hj.ht.probeScratch.groupID[batchSize-1]
	_ = hj.probeState.probeRowUnmatched[batchSize-1]
	_ = hj.probeState.buildIdx[batchSize-1]
	_ = hj.probeState.probeIdx[batchSize-1]
	_ = sel[batchSize-1]
	for i := int(0); i < batchSize; i++ {
		// Index of keys and outputs in the hash table is calculated as ID - 1.
		id := hj.ht.probeScratch.groupID[i]
		rowUnmatched := id == 0
		hj.probeState.probeRowUnmatched[i] = rowUnmatched
		if !rowUnmatched {
			hj.probeState.buildIdx[i] = int(id - 1)
		}
		{
			var __retval_0 int
			{
				{
					__retval_0 = sel[i]
				}
			}
			hj.probeState.probeIdx[i] = __retval_0
		}
	}
}

const _ = "template_distinctCollectProbeNoOuter"

func distinctCollectProbeNoOuter_false(
	hj *hashJoiner, batchSize int, nResults int, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.groupID[batchSize-1]
	_ = hj.probeState.buildIdx[batchSize-1]
	_ = hj.probeState.probeIdx[batchSize-1]
	for i := int(0); i < batchSize; i++ {
		if hj.ht.probeScratch.groupID[i] != 0 {
			// Index of keys and outputs in the hash table is calculated as ID - 1.
			hj.probeState.buildIdx[nResults] = int(hj.ht.probeScratch.groupID[i] - 1)
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

func distinctCollectProbeNoOuter_true(
	hj *hashJoiner, batchSize int, nResults int, sel []int) int {
	// Early bounds checks.
	_ = hj.ht.probeScratch.groupID[batchSize-1]
	_ = hj.probeState.buildIdx[batchSize-1]
	_ = hj.probeState.probeIdx[batchSize-1]
	_ = sel[batchSize-1]
	for i := int(0); i < batchSize; i++ {
		if hj.ht.probeScratch.groupID[i] != 0 {
			// Index of keys and outputs in the hash table is calculated as ID - 1.
			hj.probeState.buildIdx[nResults] = int(hj.ht.probeScratch.groupID[i] - 1)
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

// collect prepares the buildIdx and probeIdx arrays where the buildIdx and
// probeIdx at each index are joined to make an output row. The total number of
// resulting rows is returned.
func (hj *hashJoiner) collect(batch coldata.Batch, batchSize int, sel []int) int {
	nResults := int(0)

	if hj.spec.left.outer {
		if sel != nil {
			nResults = collectProbeOuter_true(hj, batchSize, nResults, batch, sel)
		} else {
			nResults = collectProbeOuter_false(hj, batchSize, nResults, batch, sel)
		}
	} else {
		if sel != nil {
			switch hj.spec.joinType {
			case sqlbase.LeftAntiJoin, sqlbase.ExceptAllJoin:
				nResults = collectAnti_true(hj, batchSize, nResults, batch, sel)
			default:
				nResults = collectProbeNoOuter_true(hj, batchSize, nResults, batch, sel)
			}
		} else {
			switch hj.spec.joinType {
			case sqlbase.LeftAntiJoin, sqlbase.ExceptAllJoin:
				nResults = collectAnti_false(hj, batchSize, nResults, batch, sel)
			default:
				nResults = collectProbeNoOuter_false(hj, batchSize, nResults, batch, sel)
			}
		}
	}

	return nResults
}

// distinctCollect prepares the batch with the joined output columns where the build
// row index for each probe row is given in the groupID slice. This function
// requires assumes a N-1 hash join.
func (hj *hashJoiner) distinctCollect(batch coldata.Batch, batchSize int, sel []int) int {
	nResults := int(0)

	if hj.spec.left.outer {
		nResults = batchSize

		if sel != nil {
			distinctCollectProbeOuter_true(hj, batchSize, sel)
		} else {
			distinctCollectProbeOuter_false(hj, batchSize, sel)
		}
	} else {
		if sel != nil {
			switch hj.spec.joinType {
			case sqlbase.LeftAntiJoin, sqlbase.ExceptAllJoin:
				// For LEFT ANTI and EXCEPT ALL joins we don't care whether the build
				// (right) side was distinct, so we only have single variation of COLLECT
				// method.
				nResults = collectAnti_true(hj, batchSize, nResults, batch, sel)
			default:
				nResults = distinctCollectProbeNoOuter_true(hj, batchSize, nResults, sel)
			}
		} else {
			switch hj.spec.joinType {
			case sqlbase.LeftAntiJoin, sqlbase.ExceptAllJoin:
				// For LEFT ANTI and EXCEPT ALL joins we don't care whether the build
				// (right) side was distinct, so we only have single variation of COLLECT
				// method.
				nResults = collectAnti_false(hj, batchSize, nResults, batch, sel)
			default:
				nResults = distinctCollectProbeNoOuter_false(hj, batchSize, nResults, sel)
			}
		}
	}

	return nResults
}

// execgen:inline
const _ = "template_getIdx"

// execgen:inline
const _ = "inlined_getIdx_false"

// execgen:inline
const _ = "inlined_getIdx_true"
