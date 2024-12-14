/*
Copyright © 2024 Prince kumar
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		align, _ := cmd.Flags().GetString("align")
		bold, _ := cmd.Flags().GetBool("bold")
		bgcolor, _ := cmd.Flags().GetString("bgcolor")
		fcolor, _ := cmd.Flags().GetString("fcolor")
		width, _ := cmd.Flags().GetInt("width")
		// check align var if it is valid
		if align == "left" || align == "center" || align == "right" {
			fmt.Printf("")
		} else {
			fmt.Println("\033[31;1m Please provide the valid value for align")
			os.Exit(1)
		}

		// here we have to define aling full var
		//Align := "lipgloss." + align
		var style = lipgloss.NewStyle().
			Bold(bold).
			Foreground(lipgloss.Color(bgcolor)).
			Background(lipgloss.Color(fcolor)).
			Align(lipgloss.Left).
			Width(width)
		fmt.Println(style.Render(args[0]))

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
	textCmd.Flags().StringP("align", "a", "left", "Alignment of text")
	textCmd.Flags().BoolP("bold", "b", false, "It makes text bold")
	textCmd.Flags().StringP("bgcolor", "o", " ", "Define a background colour")
	textCmd.Flags().StringP("fcolor", "f", " ", "Foreground a background colour")
	textCmd.Flags().IntP("Width", "w", 0, "Define width")

}
