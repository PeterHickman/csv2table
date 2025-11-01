package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	ep "github.com/PeterHickman/expand_path"
	"github.com/PeterHickman/toolbox"
	"io"
	"os"
	"strconv"
	"strings"
)

var delimiter = ','
var output = "unset"

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

func write_md(records [][]string) {
	for ri, record := range records {
		s := "|" + strings.Join(record, "|") + "|"
		fmt.Println(s)

		if ri == 0 {
			s := "|" + strings.Repeat("---|", len(record))
			fmt.Println(s)
		}
	}
}

func is_int(value string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func is_float(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	return true
}

func has_quotes(value string) bool {
	i := strings.Index(value, "\"")
	if i == -1 {
		return false
	}
	return true
}

func embed_quotes(value string) string {
	var new_value []byte

	new_value = append(new_value, '"')

	for i := 0; i < len(value); i++ {
		if value[i] == '"' {
			new_value = append(new_value, '\\')
		}
		new_value = append(new_value, value[i])
	}

	new_value = append(new_value, '"')

	return string(new_value)
}

func format_value(value string) string {
	lower_value := strings.ToLower(value)

	if lower_value == "true" || lower_value == "false" {
		return lower_value
	} else if lower_value == "nil" || lower_value == "null" {
		return "null"
	} else if is_int(lower_value) || is_float(lower_value) {
		return lower_value
	} else if has_quotes(value) {
		return embed_quotes(value)
	} else {
		return fmt.Sprintf("\"%s\"", value)
	}
}

func write_json(records [][]string) {
	last_record_index := len(records) - 1
	last_field_index := len(records[0]) - 1

	fmt.Println("[")

	for ri, record := range records {
		if ri != 0 {
			fmt.Println("  {")

			for i, value := range record {
				if i == last_field_index {
					fmt.Printf("    \"%s\": %s\n", records[0][i], format_value(value))
				} else {
					fmt.Printf("    \"%s\": %s,\n", records[0][i], format_value(value))
				}
			}

			if ri == last_record_index {
				fmt.Println("  }")
			} else {
				fmt.Println("  },")
			}
		}
	}

	fmt.Println("]")
}

func init() {
	d := flag.String("delimit", ",", "The character that delimit the columns")
	t := flag.Bool("table", false, "Format the output as an ascii table")
	m := flag.Bool("md", false, "Format the output as markdown")
	j := flag.Bool("json", false, "Format the output as json")

	flag.Parse()

	// Sort out the delimiter
	if *d == "\\t" {
		*d = "\t"
	}

	if len(*d) > 1 {
		dropdead(fmt.Sprintf("Column delimiter should be a single character [%s]", *d))
	}

	delimiter = []rune(*d)[0]

	if *t {
		if output != "unset" {
			dropdead(fmt.Sprintf("The format is already set to %s", output))
		} else {
			output = "table"
		}
	}

	if *m {
		if output != "unset" {
			dropdead(fmt.Sprintf("The format is already set to %s", output))
		} else {
			output = "md"
		}
	}

	if *j {
		if output != "unset" {
			dropdead(fmt.Sprintf("The format is already set to %s", output))
		} else {
			output = "json"
		}
	}

	if output == "unset" {
		output = "table"
	}
}

func main() {
	in := open_reader()

	number_of_fields, records := first_read_of_the_date(in)

	switch output {
	case "table":
		widths := field_widths(records, number_of_fields)
		write_table(records, number_of_fields, widths)
	case "md":
		write_md(records)
	case "json":
		write_json(records)
	default:
		dropdead(fmt.Sprintf("output is set to %s???", output))
	}
}
