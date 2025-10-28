package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	ep "github.com/PeterHickman/expand_path"
	"github.com/PeterHickman/toolbox"
	"io"
	"os"
	"strings"
)

var delimiter = ','

func dropdead(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func first_read_of_the_date(in io.Reader) (int, [][]string) {
	r := csv.NewReader(in)
	r.Comma = delimiter

	number_of_fields := 0
	records := [][]string{}

	// Read all the data in and find the number of fields
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			dropdead(fmt.Sprintf("%s", err))
		}

		if number_of_fields < len(record) {
			number_of_fields = len(record)
		}

		records = append(records, record)
	}

	return number_of_fields, records
}

func field_widths(records [][]string, number_of_fields int) []int {
	// Find the maximum width of each field
	widths := make([]int, number_of_fields)

	for _, record := range records {
		for i := 0; i < len(record); i++ {
			if widths[i] < len(record[i]) {
				widths[i] = len(record[i])
			}
		}
	}

	return widths
}

func open_reader() io.Reader {
	var in io.Reader

	switch len(flag.Args()) {
	case 0:
		in = os.Stdin
	case 1:
		var err error
		var filename string
		filename, err = ep.ExpandPath(flag.Arg(0))

		if err != nil {
			dropdead(fmt.Sprintf("Unable to read %s, %s\n", flag.Arg(0), err))
		}

		if !toolbox.FileExists(filename) {
			dropdead(fmt.Sprintf("Unable to read %s\n", filename))
		}

		in, err = os.Open(filename)
		if err != nil {
			dropdead(fmt.Sprintf("Unable to read %s, %s\n", flag.Arg(0), err))
		}
	default:
		dropdead("Supply one one file as an argument")
	}

	return in
}

func write_table(records [][]string, number_of_fields int, widths []int) {
	// Output the records
	for ri, record := range records {
		s := "|"
		for i := 0; i < number_of_fields; i++ {
			s += fmt.Sprintf(" %-*s ", widths[i], record[i])
			s += "|"
		}
		fmt.Println(s)

		// A seperator line after the heading
		if ri == 0 {
			s := "+"
			for i := 0; i < number_of_fields; i++ {
				s += strings.Repeat("-", widths[i]+2)
				s += "+"
			}
			fmt.Println(s)
		}
	}
}

func init() {
	d := flag.String("delimit", ",", "The character that delimit the columns")

	flag.Parse()

	if *d == "\\t" {
		*d = fmt.Sprint("\t")
	}

	if len(*d) > 1 {
		dropdead(fmt.Sprintf("Column delimiter should be a single character [%s]\n", *d))
	}

	delimiter = []rune(*d)[0]
}

func main() {
	in := open_reader()

	number_of_fields, records := first_read_of_the_date(in)

	widths := field_widths(records, number_of_fields)

	write_table(records, number_of_fields, widths)
}
