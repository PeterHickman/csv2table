package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func first_read_of_the_date(in string) (int, [][]string) {
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ','

	number_of_fields := 0
	records := [][]string{}

	// Read all the data in and find the number of fields
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
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

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	number_of_fields, records := first_read_of_the_date(in)

	widths := field_widths(records, number_of_fields)

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
