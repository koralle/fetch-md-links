/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type Options struct {
	opturl string
}

var o = &Options{}

const version = "0.1.0"

func NewCmdFmdlinks() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "fmdlinks",
		Short:   `Create Link Statement for Markdown.`,
		Example: "fmdlinks -u https://example.com",

		// $ fmdlinks --version
		// fmdlinks version 0.1.0
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {

			url := o.opturl

			doc, err := goquery.NewDocument(url)
			if err != nil {
				return err
			}
			selection := doc.Find("title")

			res := "[" + selection.Text() + "](" + url + ")"
			cmd.Println(res)
			return nil
		},
	}
	cmd.Flags().StringVarP(&o.opturl, "url", "u", "https://example.com", "url option")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewCmdFmdlinks())
}
