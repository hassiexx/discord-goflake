# Discord Goflake
[![CircleCI](https://img.shields.io/circleci/build/github/hassieswift621/discord-goflake?logo=circleci&style=flat-square)](https://circleci.com/gh/hassieswift621/discord-goflake)
[![Codacy Badge](https://img.shields.io/codacy/grade/b4770f30d53e479c9aba64cb9109a5c4?logo=codacy&style=flat-square)](https://www.codacy.com/gh/hassieswift621/discord-goflake/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hassieswift621/discord-goflake&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://img.shields.io/codacy/coverage/b4770f30d53e479c9aba64cb9109a5c4/master?logo=codacy&style=flat-square)](https://www.codacy.com/gh/hassieswift621/discord-goflake/dashboard?utm_source=github.com&utm_medium=referral&utm_content=hassieswift621/discord-goflake&utm_campaign=Badge_Coverage)
[![Godoc reference](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/hassieswift621/discord-goflake)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/hassieswift621/discord-goflake?logo=go&style=flat-square)](https://github.com/hassieswift621/discord-goflake/releases)

Discord Goflake is a simple and fast Go utility library for handling Discord snowflake IDs.

It contains utilities such as parsing snowflakes and retrieving timestamps from snowflakes.

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
```text
Snowflake to Timestamp              1000000000           0.590 ns/op
UInt64 to Snowflake                 1000000000	         0.295 ns/op
String to Snowflake          	    14054162	         86.3 ns/op
Epoch Milliseconds to Snowflake     1000000000	         0.633 ns/op
Epoch Seconds to Snowflake          71313086	         17.0 ns/op
```

## License
```text
Copyright ©2019 Hassie.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
