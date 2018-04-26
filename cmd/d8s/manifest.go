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

package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "manifest",
		Short: "Get the manifest of an image",
		Args:  cobra.ExactArgs(1),
		Run:   manifest,
	})
}

func manifest(_ *cobra.Command, args []string) {
	ref := args[0]
	i, err := getImage(ref)
	if err != nil {
		log.Fatalln(err)
	}
	manifest, err := i.RawManifest()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(string(manifest))
}
