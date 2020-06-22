/*
Copyright © 2020 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List TrueType font files",
	Long:  `List TrueType font files.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range findfont.List() {
			fontData, err := ioutil.ReadFile(p)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
			var name string
			font, err := truetype.Parse(fontData)
			if err == nil {
				name = font.Name(1)
			}
			if name == "" {
				name = "-"
			}
			if err != nil {
				fmt.Printf("name:%s\tpath:%s\terr:\"%s\"\n", name, p, err)
				continue
			}
			fmt.Printf("name:%s\tpath:%s\terr:\n", name, p)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
