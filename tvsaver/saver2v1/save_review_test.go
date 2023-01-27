// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"bytes"
	"testing"

	"github.com/spdx/tools-golang/spdx/v2_1"
)

// ===== Review section Saver tests =====
func TestSaver2_1ReviewSavesText(t *testing.T) {
	rev := &v2_1.Review{
		Reviewer:      "John Doe",
		ReviewerType:  "Person",
		ReviewDate:    "2018-10-14T10:28:00Z",
		ReviewComment: "this is a review comment",
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`Reviewer: Person: John Doe
ReviewDate: 2018-10-14T10:28:00Z
ReviewComment: this is a review comment

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderReview2_1(rev, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}

func TestSaver2_1ReviewOmitsOptionalFieldsIfEmpty(t *testing.T) {
	rev := &v2_1.Review{
		Reviewer:     "John Doe",
		ReviewerType: "Person",
		ReviewDate:   "2018-10-14T10:28:00Z",
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`Reviewer: Person: John Doe
ReviewDate: 2018-10-14T10:28:00Z

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderReview2_1(rev, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}

func TestSaver2_1ReviewWrapsMultiLine(t *testing.T) {
	rev := &v2_1.Review{
		Reviewer:     "John Doe",
		ReviewerType: "Person",
		ReviewDate:   "2018-10-14T10:28:00Z",
		ReviewComment: `this is a
multi-line review comment`,
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`Reviewer: Person: John Doe
ReviewDate: 2018-10-14T10:28:00Z
ReviewComment: <text>this is a
multi-line review comment</text>

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderReview2_1(rev, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}
