// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/sklemmer/go-autoupdate/update"
	"github.com/sklemmer/go-autoupdate/provider"
)

const Version = "0.0.1"

func init() {
	// this part is easy, simply add updater.UpdateCommand() to your rootCommand

	githubProvider := provider.NewGithubProvider(provider.NewGithubOptions("sklemmer/go-autoupdate-example"))
	updater := update.NewUpdater(githubProvider, Version)

	rootCmd.AddCommand(updater.GetUpdateCommand())
}
