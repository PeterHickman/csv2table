# csv2

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

## Alternative column names

Giving `--names tom,dick,harry` will replace the existing column names. Make sure that you have the same number of names as columns

## Data has no column names

Use `--nonames` to signal that the data has no column names and then use `--names ...` above to set them. Then you can do things like:

```bash
$ csv2 --nonames --names username,password,uid,gid,gecos,home,shell --delimit : --json /etc/passwd
[
  {
    "username": "nobody",
    "password": "*",
    "uid": -2,
    "gid": -2,
    "gecos": "Unprivileged User",
    "home": "/var/empty",
    "shell": "/usr/bin/false"
  },
  {
    "username": "root",
    "password": "*",
    "uid": 0,
    "gid": 0,
    "gecos": "System Administrator",
    "home": "/var/root",
    "shell": "/bin/sh"
  },
  {
    "username": "daemon",
    "password": "*",
    "uid": 1,
    "gid": 1,
    "gecos": "System Services",
    "home": "/var/root",
    "shell": "/usr/bin/false"
  },
  {
    "username": "_uucp",
    "password": "*",
    "uid": 4,
    "gid": 4,
    "gecos": "Unix to Unix Copy Protocol",
    "home": "/var/spool/uucp",
    "shell": "/usr/sbin/uucico"
  },
  {
    "username": "_lp",
    "password": "*",
    "uid": 26,
    "gid": 26,
    "gecos": "Printing Services",
    "home": "/var/spool/cups",
    "shell": "/usr/bin/false"
  }
]
```
