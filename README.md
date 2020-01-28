# ddboking
This CLI was created to help identify duplicated events in an Event JSON.

### Requirement
Requires ```go > 1.12``` and ``` make ``` instaled. 

### How To install
Download the zip file provided and execute ```make build``` or ```go build```. A binary named ddboking will be generated.

### How to Use
When executing ```./ddbooking``` it will show the help command output.

##### Find duplicated from text 
```bash 
./ddbooking --text='[ { "name":"Meeting", "start_date":"1985-05-12T01:05:24.311639772Z", "end_date":"1985-05-12T01:15:24.311639772Z" }, { "name":"Ocean", "start_date":"1985-05-12T01:08:24.311639772Z", "end_date":"1985-05-12T01:10:24.311639772Z" },{ "name":"Air", "start_date":"1985-05-12T01:05:24.311639772Z", "end_date":"1985-05-12T01:11:24.311639772Z" }]' import

### output expected
+---------+--------------------------------+--------------------------------+---------+--------------------------------+--------------------------------+
| EVENT 1 |           START DATE           |            END DATE            | EVENT 2 |           START DATE           |            END DATE            |
+---------+--------------------------------+--------------------------------+---------+--------------------------------+--------------------------------+
| Meeting | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:15:24.311639772  | Air     | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:11:24.311639772  |
|         | +0000 UTC                      | +0000 UTC                      |         | +0000 UTC                      | +0000 UTC                      |
| Meeting | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:15:24.311639772  | Ocean   | 1985-05-12 01:08:24.311639772  | 1985-05-12 01:10:24.311639772  |
|         | +0000 UTC                      | +0000 UTC                      |         | +0000 UTC                      | +0000 UTC                      |
| Air     | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:11:24.311639772  | Ocean   | 1985-05-12 01:08:24.311639772  | 1985-05-12 01:10:24.311639772  |
|         | +0000 UTC                      | +0000 UTC                      |         | +0000 UTC                      | +0000 UTC                      |
+---------+--------------------------------+--------------------------------+---------+--------------------------------+--------------------------------+


```

##### Find duplicated from Json file
```bash 
### Example of an Event json file https://gist.github.com/lsantanna87/a7acf127f138cbc487adcd361a674500
 ./ddbooking --file=<path_to_filename>.json import
 
 ## output expected
+---------+--------------------------------+--------------------------------+---------+--------------------------------+--------------------------------+
| EVENT 1 |           START DATE           |            END DATE            | EVENT 2 |           START DATE           |            END DATE            |
+---------+--------------------------------+--------------------------------+---------+--------------------------------+--------------------------------+
| Meeting | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:15:24.311639772  | Air     | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:11:24.311639772  |
|         | +0000 UTC                      | +0000 UTC                      |         | +0000 UTC                      | +0000 UTC                      |
| Meeting | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:15:24.311639772  | Ocean   | 1985-05-12 01:08:24.311639772  | 1985-05-12 01:10:24.311639772  |
|         | +0000 UTC                      | +0000 UTC                      |         | +0000 UTC                      | +0000 UTC                      |
| Air     | 1985-05-12 01:05:24.311639772  | 1985-05-12 01:11:24.311639772  | Ocean   | 1985-05-12 01:08:24.311639772  | 1985-05-12 01:10:24.311639772  |
|         | +0000 UTC                      | +0000 UTC                      |         | +0000 UTC                      | +0000 UTC                      |
+---------+--------------------------------+--------------------------------+---------+--------------------------------+--------------------------------+

```
##### Validate if Events from text are valid
```bash 
./ddbooking --text='[ { "name":"Meeting", "start_date":"1985-05-12T01:05:24.311639772Z", "end_date":"1985-05-12T01:15:24.311639772Z" }, { "name":"Ocean", "start_date":"1985-05-12T01:08:24.311639772Z", "end_date":"1985-05-12T01:10:24.311639772Z" },{ "name":"Air", "start_date":"1985-05-12T01:05:24.311639772Z", "end_date":"1985-05-12T01:11:24.311639772Z" }]' validate

### output expected
+-------------------+
| ARE EVENTS VALID? |
+-------------------+
| true              |
+-------------------+
```

##### Validate if Events from Json are valid
```bash 
### Example of an Event json file https://gist.github.com/lsantanna87/a7acf127f138cbc487adcd361a674500
 ./ddbooking --file=<path_to_filename>.json validate
 
 ## output expected
+-------------------+
| ARE EVENTS VALID? |
+-------------------+
| true              |
+-------------------+
```

##### Helper 
``` bash ./ddbooking --help```
```bash
NAME:
   Search for Events Overlapping - A new CLI application

USAGE:
   ddbooking [global options] command [command options] [arguments...]

COMMANDS:
   import    Import Events
   validate  Validate Events
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file FILE   Load Events from json FILE eg: https://gist.github.com/lsantanna87/a7acf127f138cbc487adcd361a674500
   --text value  Load Events from text in JSON format eg: https://gist.github.com/lsantanna87/5aeb75a0e9affc2eb0cfc8f087acb4da
   --help, -h    show help (default: false)
```

### Other options for the Makefile
```make 
## Makefile options 
build                          Build ddbooking Binary
ci                             Simulates CI.
cleanup                        Deletes Temp Files.
lint                           Runs linter
test                           Runs tests
test-with-coverage             Runs tests with coverage
```

### TODO
* Consider others date format

