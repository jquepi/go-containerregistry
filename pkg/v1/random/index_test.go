// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package random

import (
	"testing"

	"github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/google/go-containerregistry/pkg/v1/validate"
)

func TestRandomIndex(t *testing.T) {
	ii, err := Index(1024, 5, 3)
	if err != nil {
		t.Fatalf("Error loading index: %v", err)
	}

	if err := validate.Index(ii); err != nil {
		t.Errorf("validate.Index() = %v", err)
	}

	digest, err := ii.Digest()
	if err != nil {
		t.Fatalf("Digest(): unexpected err: %v", err)
	}

	if _, err := ii.Image(digest); err == nil {
		t.Errorf("Image(%s): expected err, got nil", digest)
	}

	if _, err := ii.ImageIndex(digest); err == nil {
		t.Errorf("ImageIndex(%s): expected err, got nil", digest)
	}

	mt, err := ii.MediaType()
	if err != nil {
		t.Errorf("MediaType(): unexpected err: %v", err)
	}

	if got, want := mt, types.OCIImageIndex; got != want {
		t.Errorf("MediaType(): got: %v, want: %v", got, want)
	}

	man, err := ii.IndexManifest()
	if err != nil {
		t.Errorf("IndexManifest(): unexpected err: %v", err)
	}

	if got, want := man.MediaType, types.OCIImageIndex; got != want {
		t.Errorf("MediaType: got: %v, want: %v", got, want)
	}
}
