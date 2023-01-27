// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package spdxlib

import (
	"fmt"

	"github.com/spdx/tools-golang/spdx/common"
	"github.com/spdx/tools-golang/spdx/v2_1"
	"github.com/spdx/tools-golang/spdx/v2_2"
	"github.com/spdx/tools-golang/spdx/v2_3"
)

// ValidateDocument2_1 returns an error if the Document is found to be invalid, or nil if the Document is valid.
// Currently, this only verifies that all Element IDs mentioned in Relationships exist in the Document as either a
// Package or an UnpackagedFile.
func ValidateDocument2_1(doc *v2_1.Document) error {
	// cache a map of valid package IDs for quick lookups
	validElementIDs := make(map[common.ElementID]bool)
	for _, docPackage := range doc.Packages {
		validElementIDs[docPackage.PackageSPDXIdentifier] = true
	}

	for _, unpackagedFile := range doc.Files {
		validElementIDs[unpackagedFile.FileSPDXIdentifier] = true
	}

	// add the Document element ID
	validElementIDs[common.MakeDocElementID("", "DOCUMENT").ElementRefID] = true

	for _, relationship := range doc.Relationships {
		if !validElementIDs[relationship.RefA.ElementRefID] {
			return fmt.Errorf("%s used in relationship but no such package exists", string(relationship.RefA.ElementRefID))
		}

		if !validElementIDs[relationship.RefB.ElementRefID] {
			return fmt.Errorf("%s used in relationship but no such package exists", string(relationship.RefB.ElementRefID))
		}
	}

	return nil
}

// ValidateDocument2_2 returns an error if the Document is found to be invalid, or nil if the Document is valid.
// Currently, this only verifies that all Element IDs mentioned in Relationships exist in the Document as either a
// Package or an UnpackagedFile.
func ValidateDocument2_2(doc *v2_2.Document) error {
	// cache a map of package IDs for quick lookups
	validElementIDs := make(map[common.ElementID]bool)
	for _, docPackage := range doc.Packages {
		validElementIDs[docPackage.PackageSPDXIdentifier] = true
	}

	for _, unpackagedFile := range doc.Files {
		validElementIDs[unpackagedFile.FileSPDXIdentifier] = true
	}

	// add the Document element ID
	validElementIDs[common.MakeDocElementID("", "DOCUMENT").ElementRefID] = true

	for _, relationship := range doc.Relationships {
		if !validElementIDs[relationship.RefA.ElementRefID] {
			return fmt.Errorf("%s used in relationship but no such package exists", string(relationship.RefA.ElementRefID))
		}

		if !validElementIDs[relationship.RefB.ElementRefID] {
			return fmt.Errorf("%s used in relationship but no such package exists", string(relationship.RefB.ElementRefID))
		}
	}

	return nil
}

// ValidateDocument2_3 returns an error if the Document is found to be invalid, or nil if the Document is valid.
// Currently, this only verifies that all Element IDs mentioned in Relationships exist in the Document as either a
// Package or an UnpackagedFile.
func ValidateDocument2_3(doc *v2_3.Document) error {
	// cache a map of package IDs for quick lookups
	validElementIDs := make(map[common.ElementID]bool)
	for _, docPackage := range doc.Packages {
		validElementIDs[docPackage.PackageSPDXIdentifier] = true
	}

	for _, unpackagedFile := range doc.Files {
		validElementIDs[unpackagedFile.FileSPDXIdentifier] = true
	}

	// add the Document element ID
	validElementIDs[common.MakeDocElementID("", "DOCUMENT").ElementRefID] = true

	for _, relationship := range doc.Relationships {
		if !validElementIDs[relationship.RefA.ElementRefID] {
			return fmt.Errorf("%s used in relationship but no such package exists", string(relationship.RefA.ElementRefID))
		}

		if !validElementIDs[relationship.RefB.ElementRefID] {
			return fmt.Errorf("%s used in relationship but no such package exists", string(relationship.RefB.ElementRefID))
		}
	}

	return nil
}
