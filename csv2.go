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
var names []string
var nonames = false

func dropdead(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func first_read_of_the_date(in io.Reader) [][]string {
	r := csv.NewReader(in)
	r.Comma = delimiter
	r.Comment = '#'

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

		records = append(records, record)
	}

	return records
}

func field_widths(records [][]string) []int {
	widths := make([]int, len(records[0]))

	if len(names) > 0 {
		for i := 0; i < len(names); i++ {
			if widths[i] < len(names[i]) {
				widths[i] = len(names[i])
			}
		}
	}

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

func table_line(record []string, widths []int) {
	s := "|"
	for i := 0; i < len(record); i++ {
		s += fmt.Sprintf(" %-*s ", widths[i], record[i])
		s += "|"
	}
	fmt.Println(s)
}

func table_header(widths []int) {
	s := "+"
	for i := 0; i < len(widths); i++ {
		s += strings.Repeat("-", widths[i]+2)
		s += "+"
	}
	fmt.Println(s)
}

func are_names_defined(record_size int) bool {
	if len(names) > 0 {
		if len(names) == record_size {
			return true
		} else {
			dropdead(fmt.Sprintf("The --name option defined %d columns, the data has %d", len(names), record_size))
		}
	}

	return false
}

func write_table(records [][]string) {
	widths := field_widths(records)

	use_names := are_names_defined(len(widths))

	for ri, record := range records {
		if ri == 0 {
			if use_names {
				table_line(names, widths)
				table_header(widths)
				table_line(record, widths)
			} else {
				table_line(record, widths)
				table_header(widths)
			}
		} else {
			table_line(record, widths)
		}
	}
}

func md_line(record []string) {
	s := "|" + strings.Join(record, "|") + "|"
	fmt.Println(s)
}

func md_headers(record []string) {
	s := "|" + strings.Repeat("---|", len(record))
	fmt.Println(s)
}

func write_md(records [][]string) {
	use_names := are_names_defined(len(records[0]))

	for ri, record := range records {
		if ri == 0 {
			if use_names {
				md_line(names)
				md_headers(record)
				md_line(record)
			} else {
				md_line(record)
				md_headers(record)
			}
		} else {
			md_line(record)
		}
	}
}

func is_int(value string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
	return err == nil
}

func is_float(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func has_quotes(value string) bool {
	i := strings.Index(value, "\"")
	return i != -1
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
	are_names_defined(len(records[0]))

	last_record_index := len(records) - 1
	last_field_index := len(records[0]) - 1

	var headers []string
	first_row := 1

	if nonames {
		headers = names
		first_row = 0
	} else if len(names) > 0 {
		headers = names
	} else {
		headers = records[0]
	}

	fmt.Println("[")

	for ri, record := range records {
		if ri >= first_row {
			fmt.Println("  {")

			for i, value := range record {
				if i == last_field_index {
					fmt.Printf("    \"%s\": %s\n", headers[i], format_value(value))
				} else {
					fmt.Printf("    \"%s\": %s,\n", headers[i], format_value(value))
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
	n := flag.String("names", "", "Optional column headers when the file has none. Seperate with commas")
	x := flag.Bool("nonames", false, "When the csv file does not have column names, use with --names")

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

	if *n != "" {
		names = strings.Split(*n, ",")
	}

	if *x {
		if len(names) == 0 {
			dropdead("File has no column names, via --nonames, but no names supplied, via --names")
		} else {
			nonames = true
		}
	}
}

func main() {
	in := open_reader()

	records := first_read_of_the_date(in)

	switch output {
	case "table":
		write_table(records)
	case "md":
		write_md(records)
	case "json":
		write_json(records)
	default:
		dropdead(fmt.Sprintf("output is set to %s???", output))
	}
}
