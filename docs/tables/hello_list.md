# Table: hello_list

This table defines a List hydrate.

```
func tableHelloList(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_list",
		Description: "hello using List/ListConfig",
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: helloCols(),
	}
}
```

## Examples

These examples use `listGreeting`. 

### 1

```
select id, greeting, json, jsonb_pretty(_ctx) from hello_list
```

```
+----+----------+-------------------+-----------------------------------------------------+
| id | greeting | json              | jsonb_pretty                                        |
+----+----------+-------------------+-----------------------------------------------------+
| 3  | Hello    | {"hello":"world"} | {                                                   |
|    |          |                   |     "steampipe": {                                  |
|    |          |                   |         "sdk_version": "5.8.0"                      |
|    |          |                   |     },                                              |
|    |          |                   |     "diagnostics": {                                |
|    |          |                   |         "calls": [                                  |
|    |          |                   |             {                                       |
|    |          |                   |                 "type": "list",                     |
|    |          |                   |                 "scope_values": {                   |
|    |          |                   |                     "table": "hello_list",          |
|    |          |                   |                     "connection": "hello",          |
|    |          |                   |                     "function_name": "listGreeting" |
|    |          |                   |                 },                                  |
|    |          |                   |                 "function_name": "listGreeting",    |
|    |          |                   |                 "rate_limiters": [                  |
|    |          |                   |                 ],                                  |
|    |          |                   |                 "rate_limiter_delay_ms": 0          |
|    |          |                   |             }                                       |
|    |          |                   |         ]                                           |
|    |          |                   |     },                                              |
|    |          |                   |     "connection_name": "hello"                      |
|    |          |                   | }                                                   |
| 2  | Hello    | {"hello":"world"} | {                                                   |
... more data ...
```

### 2

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered by Postgres down to 1 row.

```
select id, greeting, json->>'hello' as json_value, jsonb_pretty(_ctx) as _ctx from hello_list where id = 2
```

```
+----+----------+------------+-----------------------------------------------------+
| id | greeting | json_value | _ctx                                                |
+----+----------+------------+-----------------------------------------------------+
| 2  | Hello    | world      | {                                                   |
|    |          |            |     "steampipe": {                                  |
|    |          |            |         "sdk_version": "5.8.0"                      |
|    |          |            |     },                                              |
|    |          |            |     "diagnostics": {                                |
|    |          |            |         "calls": [                                  |
|    |          |            |             {                                       |
|    |          |            |                 "type": "list",                     |
|    |          |            |                 "scope_values": {                   |
|    |          |            |                     "table": "hello_list",          |
|    |          |            |                     "connection": "hello",          |
|    |          |            |                     "function_name": "listGreeting" |
|    |          |            |                 },                                  |
|    |          |            |                 "function_name": "listGreeting",    |
|    |          |            |                 "rate_limiters": [                  |
|    |          |            |                 ],                                  |
|    |          |            |                 "rate_limiter_delay_ms": 0          |
|    |          |            |             }                                       |
|    |          |            |         ]                                           |
|    |          |            |     },                                              |
|    |          |            |     "connection_name": "hello"                      |
|    |          |            | }                                                   |
+----+----------+------------+-----------------------------------------------------+
```

### 3

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered down to 0 rows.

```
select id, greeting, json->>'hello' as json_value, jsonb_pretty(_ctx) as _ctx from hello_list where id = 17

```

```
+----+----------+------------+------+
| id | greeting | json_value | _ctx |
+----+----------+------------+------+
+----+----------+------------+------+
```

### 4

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered down to 2 rows.

```
select id, greeting, json->>'hello' as json_value, jsonb_pretty(_ctx) as _ctx from hello_list where id in (1,2,17)
```

```
+----+----------+------------+-----------------------------------------------------+
| id | greeting | json_value | _ctx                                                |
+----+----------+------------+-----------------------------------------------------+
| 1  | Hello    | world      | {                                                   |
|    |          |            |     "steampipe": {                                  |
|    |          |            |         "sdk_version": "5.8.0"                      |
|    |          |            |     },                                              |
|    |          |            |     "diagnostics": {                                |
|    |          |            |         "calls": [                                  |
|    |          |            |             {                                       |
|    |          |            |                 "type": "list",                     |
|    |          |            |                 "scope_values": {                   |
|    |          |            |                     "table": "hello_list",          |
|    |          |            |                     "connection": "hello",          |
|    |          |            |                     "function_name": "listGreeting" |
|    |          |            |                 },                                  |
|    |          |            |                 "function_name": "listGreeting",    |
|    |          |            |                 "rate_limiters": [                  |
|    |          |            |                 ],                                  |
|    |          |            |                 "rate_limiter_delay_ms": 0          |
|    |          |            |             }                                       |
|    |          |            |         ]                                           |
|    |          |            |     },                                              |
|    |          |            |     "connection_name": "hello"                      |
|    |          |            | }                                                   |
| 2  | Hello    | world      | {                                                   |
|    |          |            |     "steampipe": {                                  |
|    |          |            |         "sdk_version": "5.8.0"                      |
|    |          |            |     },                                              |
|    |          |            |     "diagnostics": {                                |
|    |          |            |         "calls": [                                  |
|    |          |            |             {                                       |
|    |          |            |                 "type": "list",                     |
|    |          |            |                 "scope_values": {                   |
|    |          |            |                     "table": "hello_list",          |
|    |          |            |                     "connection": "hello",          |
|    |          |            |                     "function_name": "listGreeting" |
|    |          |            |                 },                                  |
|    |          |            |                 "function_name": "listGreeting",    |
|    |          |            |                 "rate_limiters": [                  |
|    |          |            |                 ],                                  |
|    |          |            |                 "rate_limiter_delay_ms": 0          |
|    |          |            |             }                                       |
|    |          |            |         ]                                           |
|    |          |            |     },                                              |
|    |          |            |     "connection_name": "hello"                      |
|    |          |            | }                                                   |
+----+----------+------------+-----------------------------------------------------+
```

