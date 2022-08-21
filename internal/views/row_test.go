package views

import (
	"fmt"
	rw "github.com/mattn/go-runewidth"
	"testing"
	"unicode/utf8"
)

func TestRows_SplitLine(t *testing.T) {
	s := `func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	debugP = rootCmd.Flags().BoolP("debug", "d", true, "sets log level to debug")
}`
	rows, err := NewRows([]byte(s))
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	rows.SplitLine(0, 4)

	fmt.Println(rows)
}

func Test_Rune(t *testing.T) {
	fmt.Println(utf8.DecodeRune([]byte("")))

	fmt.Println(rw.RuneWidth(utf8.RuneError))

}
