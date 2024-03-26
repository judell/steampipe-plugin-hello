# Table: hello_get_or_list

This table defines both Get and List hydrates.

```
func tableHelloGetOrList(ctx context.Context) *plugin.Table { 
	return &plugin.Table{
		Name:        "hello_get_or_list",
		Description: "hello using Get/GetConfig or List/ListConfig",
		Get: &plugin.GetConfig{
			Hydrate: getGreeting,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: helloCols(),
	}
}
```
## Examples

These examples show different ways Steampipe can use `getGreeting` or `listGreeting` to satisfy a query.

### 1 uses listGreeting

```
select id, greeting, json, jsonb_pretty(_ctx) from hello_get_or_list

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
|    |          |                   |                     "table": "hello_get_or_list",   |
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
| 1  | Hello    | {"hello":"world"} | {                                                   |
... more data ...
select id, greeting, json, jsonb_pretty(_ctx) from hello_get_or_list
+----+----------+-------------------+-----------------------------------------------------+
```

### 2 uses listGreeting

`listGreeting` runs to completion, returning ids 1, 2, and 3, then they're filtered down to 2 by Postgres

```
select id, greeting, json, jsonb_pretty(_ctx) from hello_get_or_list where id = '2'
```

```
+----+----------+-------------------+-----------------------------------------------------+
| id | greeting | json              | jsonb_pretty                                        |
+----+----------+-------------------+-----------------------------------------------------+
| 2  | Hello    | {"hello":"world"} | {                                                   |
|    |          |                   |     "steampipe": {                                  |
|    |          |                   |         "sdk_version": "5.8.0"                      |
|    |          |                   |     },                                              |
|    |          |                   |     "diagnostics": {                                |
|    |          |                   |         "calls": [                                  |
|    |          |                   |             {                                       |
|    |          |                   |                 "type": "list",                     |
|    |          |                   |                 "scope_values": {                   |
|    |          |                   |                     "table": "hello_get_or_list",   |
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
+----+----------+-------------------+-----------------------------------------------------+
```


### 3 uses getGreeting

```
select id, greeting, json, jsonb_pretty(_ctx) from hello_get_or_list where id in (1,2,17)
```

```
+----+----------+-------------------+----------------------------------------------------+
| id | greeting | json              | jsonb_pretty                                       |
+----+----------+-------------------+----------------------------------------------------+
| 1  | Hello    | {"hello":"world"} | {                                                  |
|    |          |                   |     "steampipe": {                                 |
|    |          |                   |         "sdk_version": "5.8.0"                     |
|    |          |                   |     },                                             |
|    |          |                   |     "diagnostics": {                               |
|    |          |                   |         "calls": [                                 |
|    |          |                   |             {                                      |
|    |          |                   |                 "type": "get",                     |
|    |          |                   |                 "scope_values": {                  |
|    |          |                   |                     "table": "hello_get_or_list",  |
|    |          |                   |                     "connection": "hello",         |
|    |          |                   |                     "function_name": "getGreeting" |
|    |          |                   |                 },                                 |
|    |          |                   |                 "function_name": "getGreeting",    |
|    |          |                   |                 "rate_limiters": [                 |
|    |          |                   |                 ],                                 |
|    |          |                   |                 "rate_limiter_delay_ms": 0         |
|    |          |                   |             }                                      |
|    |          |                   |         ]                                          |
|    |          |                   |     },                                             |
|    |          |                   |     "connection_name": "hello"                     |
|    |          |                   | }                                                  |
| 17 | Hello    | {"hello":"world"} | {                                                  |
... more data ...
+----+----------+-------------------+----------------------------------------------------+
```


### 4 uses listGreeting

`listGreeting` runs to completion, returning ids 1, 2, and 3, then they're filtered down to 1, 2 by Postgres

```
with ids as ( select 1 as id union select 2 union select 17 )
select id, greeting, json, jsonb_pretty(_ctx)
from 
  hello_get_or_list
join 
  ids
using
  (id)
```

```
+----+----------+-------------------+-----------------------------------------------------+
| id | greeting | json              | jsonb_pretty                                        |
+----+----------+-------------------+-----------------------------------------------------+
| 1  | Hello    | {"hello":"world"} | {                                                   |
|    |          |                   |     "steampipe": {                                  |
|    |          |                   |         "sdk_version": "5.8.0"                      |
|    |          |                   |     },                                              |
|    |          |                   |     "diagnostics": {                                |
|    |          |                   |         "calls": [                                  |
|    |          |                   |             {                                       |
|    |          |                   |                 "type": "list",                     |
|    |          |                   |                 "scope_values": {                   |
|    |          |                   |                     "table": "hello_get_or_list",   |
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
|    |          |                   |     "steampipe": {                                  |
|    |          |                   |         "sdk_version": "5.8.0"                      |
|    |          |                   |     },                                              |
|    |          |                   |     "diagnostics": {                                |
|    |          |                   |         "calls": [                                  |
|    |          |                   |             {                                       |
|    |          |                   |                 "type": "list",                     |
|    |          |                   |                 "scope_values": {                   |
|    |          |                   |                     "table": "hello_get_or_list",   |
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
+----+----------+-------------------+-----------------------------------------------------+
```

