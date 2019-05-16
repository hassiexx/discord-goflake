# Discord Goflake [![CircleCI](https://circleci.com/gh/hassieswift621/discord-goflake/tree/master.svg?style=svg)](https://circleci.com/gh/hassieswift621/discord-goflake/tree/master) [![Documentation](https://godoc.org/github.com/hassieswift621/discord-goflake?status.svg)](http://godoc.org/github.com/hassieswift621/discord-goflake) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/13bd27ca8f1f41ee9cad4c7040ea127a)](https://www.codacy.com/app/hassieswift621/discord-goflake?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hassieswift621/discord-goflake&amp;utm_campaign=Badge_Grade) [![Codacy Badge](https://api.codacy.com/project/badge/Coverage/13bd27ca8f1f41ee9cad4c7040ea127a)](https://www.codacy.com/app/hassieswift621/discord-goflake?utm_source=github.com&utm_medium=referral&utm_content=hassieswift621/discord-goflake&utm_campaign=Badge_Coverage)
Discord Goflake is a simple Go utility library for handling Discord snowflake IDs.

It contains utilities such as parsing snowflakes and retrieving timestamps from snowflakes.

The current version is 1.0.0.

## Tutorial
**Retrieving timestamp from a raw string ID**
```go
snowflake, err := dgoflake.ParseString("577685297450582016")
if err == nil {
	timestamp := snowflake.Timestamp()
}
```

**Converting epoch milliseconds timestamp to snowflake**
```go
snowflake, err := dgoflake.ParseEpochMilli(1557801521446)
```

See the godocs for the full reference.

## Benchmarks
```
Snowflake to Timestamp          2000000000	        0.43 ns/op
UInt64 to Snowflake             2000000000	        0.33 ns/op
String to Snowflake             30000000	        46.8 ns/op
Epoch Milli to Snowflake        2000000000	        0.31 ns/op
Epoch Sec to Snowflake          100000000	        17.4 ns/op
```

## License
Copyright &copy;2019 Hassie.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
