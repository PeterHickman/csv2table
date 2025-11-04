# csv2

## TODO
1. `csv2 --delimit : /etc/passwd` we need to pass in headers when the file has none

Convert a csv file into either a nice ascii table, a markdown table or json

We are assuming that the first line of the file is the column names. The default the delimiter is a `,` but if you need to change it pass `--delimit '\t'` to use a tab (for example)

Can be used as a pipe too


## `csv2 --table data.csv` or `csv2 data.csv` as table is the default

```
  | first_name | last_name | username |
  +------------+-----------+----------+
  | Rob        | Pike      | rob      |
  | Ken        | Thompson  | ken      |
  | Robert     | Griesemer | gri      |
```

## `csv2 --md data.csv`

```markdown
  |first_name|last_name|username|
  |---|---|---|
  |Rob|Pike|rob|
  |Ken|Thompson|ken|
  |Robert|Griesemer|gri|
```

## `csv2 --json data.csv`

```json
  [
    {
      "first_name": "Rob",
      "last_name": "Pike",
      "username": "rob"
    },
    {
      "first_name": "Ken",
      "last_name": "Thompson",
      "username": "ken"
    },
    {
      "first_name": "Robert",
      "last_name": "Griesemer",
      "username": "gri"
    }
  ]
```
