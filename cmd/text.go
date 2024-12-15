/*
Copyright © 2024 Prince kumar
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "This is used to print some text usign some advance options",
	Long:  `Using this commnad you can print text with some rich functionality.`,
	Run: func(cmd *cobra.Command, args []string) {
		align, _ := cmd.Flags().GetString("align")
		bold, _ := cmd.Flags().GetBool("bold")
		bgcolor, _ := cmd.Flags().GetString("bgcolor")
		fcolor, _ := cmd.Flags().GetString("fcolor")
		width, _ := cmd.Flags().GetInt("width")
		ptop, _ := cmd.Flags().GetInt("ptop")
		pleft, _ := cmd.Flags().GetInt("pleft")
		// define your style here
		var style = lipgloss.NewStyle().
			Bold(bold).
			Foreground(lipgloss.Color(bgcolor)).
			Background(lipgloss.Color(fcolor)).
			Align(lipgloss.Center).
			PaddingTop(ptop).
			PaddingLeft(pleft).
			Width(width)
		// check align var if it is valid
		if align == "center" {
			// inherit the property of style
			var stylec = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Inherit(style)
			fmt.Println(stylec.Render(args[0]))
		} else if align == "right" {
			// inherit the property of style
			var styler = lipgloss.NewStyle().
				Align(lipgloss.Right).
				Inherit(style)
			fmt.Println(styler.Render(args[0]))
		} else if align == "left" {
			// inherit the property of style
			var stylel = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Inherit(style)
			fmt.Println(stylel.Render(args[0]))
		} else {
			fmt.Println(style.Render(args[0]))
		}
		// here we have to define aling full var
		//Align := "lipgloss." + align

	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Add subcommands here
	textCmd.Flags().StringP("align", "a", "", "Alignment of text")
	textCmd.Flags().BoolP("bold", "b", false, "It makes text bold")
	textCmd.Flags().StringP("bgcolor", "o", " ", "Define a background colour")
	textCmd.Flags().StringP("fcolor", "f", " ", "Foreground a background colour")
	textCmd.Flags().IntP("width", "w", 0, "Width of text ")
	textCmd.Flags().IntP("ptop", "t", 0, "Padding top")
	textCmd.Flags().IntP("pleft", "p", 0, "Padding left")

}
