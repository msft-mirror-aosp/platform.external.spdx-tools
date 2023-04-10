// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package parser2v1

import (
	"testing"

	"github.com/spdx/tools-golang/spdx/common"
)

// ===== Helper function tests =====

func TestCanExtractSubvalues(t *testing.T) {
	subkey, subvalue, err := extractSubs("SHA1: abc123")
	if err != nil {
		t.Errorf("got error when calling extractSubs: %v", err)
	}
	if subkey != "SHA1" {
		t.Errorf("got %v for subkey", subkey)
	}
	if subvalue != "abc123" {
		t.Errorf("got %v for subvalue", subvalue)
	}
}

func TestReturnsErrorForInvalidSubvalueFormat(t *testing.T) {
	_, _, err := extractSubs("blah")
	if err == nil {
		t.Errorf("expected error when calling extractSubs for invalid format (0 colons), got nil")
	}
}

func TestCanExtractDocumentAndElementRefsFromID(t *testing.T) {
	// test with valid ID in this document
	helperForExtractDocElementID(t, "SPDXRef-file1", false, "", "file1")
	// test with valid ID in another document
	helperForExtractDocElementID(t, "DocumentRef-doc2:SPDXRef-file2", false, "doc2", "file2")
	// test with invalid ID in this document
	helperForExtractDocElementID(t, "a:SPDXRef-file1", true, "", "")
	helperForExtractDocElementID(t, "file1", true, "", "")
	helperForExtractDocElementID(t, "SPDXRef-", true, "", "")
	helperForExtractDocElementID(t, "SPDXRef-file1:", true, "", "")
	// test with invalid ID in another document
	helperForExtractDocElementID(t, "DocumentRef-doc2", true, "", "")
	helperForExtractDocElementID(t, "DocumentRef-doc2:", true, "", "")
	helperForExtractDocElementID(t, "DocumentRef-doc2:SPDXRef-", true, "", "")
	helperForExtractDocElementID(t, "DocumentRef-doc2:a", true, "", "")
	helperForExtractDocElementID(t, "DocumentRef-:", true, "", "")
	helperForExtractDocElementID(t, "DocumentRef-:SPDXRef-file1", true, "", "")
	// test with invalid formats
	helperForExtractDocElementID(t, "DocumentRef-doc2:SPDXRef-file1:file2", true, "", "")
}

func helperForExtractDocElementID(t *testing.T, tst string, wantErr bool, wantDoc string, wantElt string) {
	deID, err := extractDocElementID(tst)
	if err != nil && wantErr == false {
		t.Errorf("testing %v: expected nil error, got %v", tst, err)
	}
	if err == nil && wantErr == true {
		t.Errorf("testing %v: expected non-nil error, got nil", tst)
	}
	if deID.DocumentRefID != wantDoc {
		if wantDoc == "" {
			t.Errorf("testing %v: want empty string for DocumentRefID, got %v", tst, deID.DocumentRefID)
		} else {
			t.Errorf("testing %v: want %v for DocumentRefID, got %v", tst, wantDoc, deID.DocumentRefID)
		}
	}
	if deID.ElementRefID != common.ElementID(wantElt) {
		if wantElt == "" {
			t.Errorf("testing %v: want emptyString for ElementRefID, got %v", tst, deID.ElementRefID)
		} else {
			t.Errorf("testing %v: want %v for ElementRefID, got %v", tst, wantElt, deID.ElementRefID)
		}
	}
}

func TestCanExtractElementRefsOnlyFromID(t *testing.T) {
	// test with valid ID in this document
	helperForExtractElementID(t, "SPDXRef-file1", false, "file1")
	// test with valid ID in another document
	helperForExtractElementID(t, "DocumentRef-doc2:SPDXRef-file2", true, "")
	// test with invalid ID in this document
	helperForExtractElementID(t, "a:SPDXRef-file1", true, "")
	helperForExtractElementID(t, "file1", true, "")
	helperForExtractElementID(t, "SPDXRef-", true, "")
	helperForExtractElementID(t, "SPDXRef-file1:", true, "")
	// test with invalid ID in another document
	helperForExtractElementID(t, "DocumentRef-doc2", true, "")
	helperForExtractElementID(t, "DocumentRef-doc2:", true, "")
	helperForExtractElementID(t, "DocumentRef-doc2:SPDXRef-", true, "")
	helperForExtractElementID(t, "DocumentRef-doc2:a", true, "")
	helperForExtractElementID(t, "DocumentRef-:", true, "")
	helperForExtractElementID(t, "DocumentRef-:SPDXRef-file1", true, "")
}

func helperForExtractElementID(t *testing.T, tst string, wantErr bool, wantElt string) {
	eID, err := extractElementID(tst)
	if err != nil && wantErr == false {
		t.Errorf("testing %v: expected nil error, got %v", tst, err)
	}
	if err == nil && wantErr == true {
		t.Errorf("testing %v: expected non-nil error, got nil", tst)
	}
	if eID != common.ElementID(wantElt) {
		if wantElt == "" {
			t.Errorf("testing %v: want emptyString for ElementRefID, got %v", tst, eID)
		} else {
			t.Errorf("testing %v: want %v for ElementRefID, got %v", tst, wantElt, eID)
		}
	}
}
