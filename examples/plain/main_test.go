package main

import (
	"testing"

	"inhouse"
)

func TestInitExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("init", false)

	if err != nil {
		t.Error(err)
		return
	}

	if got == nil {
		t.Error("check should not be empty")
		return
	}

	if !got.Contained {
		t.Error("init function should exist")
	}
}

func TestInternalFuctionExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("internalFunction", false)

	if err != nil {
		t.Error(err)
		return
	}

	if got == nil {
		t.Error("check should not be empty")
		return
	}

	if !got.Contained {
		t.Error("internalFunction function should exist")
	}
}

func TestInternalFunctionNotExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("prohibidedFunction", false)

	if err != nil {
		t.Error(err)
		return
	}

	if got == nil {
		t.Error("check should not be empty")
		return
	}

	if got.Contained {
		t.Error("prohibidedFunction function should not be present")
	}
}

func TestEnsureExportedFunctionExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("ExportedFunction", false)

	if err != nil {
		t.Error(err)
		return
	}

	if got == nil {
		t.Error("check should not be empty")
		return
	}

	if !got.Contained {
		t.Error("ExportedFunction should exist")
	}
}
