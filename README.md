Timetasker CLI
==============

Provides a nice command line interface to submit time sheets using [Time
Task](https://timetask.com). This command will submit a single time entry.

The project improves upon a [Chrome
extension](https://github.com/teddywing/chrome-timetasker) that auto-fills the
time sheet form on the website.


## Usage
This will submit a time entry for the "example" project on the current day with
a duration of 7 hours and an empty description:

	$ timetasker --project example

Here we set a custom time of 4.5 hours:

	$ timetasker --project example --time 4.5

Now we specify a date and add a description:

	$ timetasker --project example --date 2017-05-31 --description "Worked on Timetasker"


## License
Copyright © 2017 Teddy Wing. Licensed under the GNU GPLv3+ (see the included
COPYING file).
