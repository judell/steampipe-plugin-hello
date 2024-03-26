# Table: hello_get

This tables defines a Get hydrate.

```
func tableHelloGet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_get",
		Description: "hello using Get/GetConfig",
		Get: &plugin.GetConfig{
			Hydrate: getGreeting,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: helloCols(),
	}
}
```

## Examples

These examples use `getGreeting`.

### 1

```
select id, greeting, json, jsonb_pretty(_ctx) from hello_get
```

```
Error: rpc error: code = Internal desc = 'Get' call for table 'hello_get' is missing 1 required qual: column:'id' operator: = (SQLSTATE HV000)
```

### 2

```
select id, greeting, json, jsonb_pretty(_ctx) as _ctx from hello_get where id = '17'
```

```
+----+----------+-------------------+----------------------------------------------------+
| id | greeting | json              | _ctx                                               |
+----+----------+-------------------+----------------------------------------------------+
| 17 | Hello    | {"hello":"world"} | {                                                  |
|    |          |                   |     "steampipe": {                                 |
|    |          |                   |         "sdk_version": "5.8.0"                     |
|    |          |                   |     },                                             |
|    |          |                   |     "diagnostics": {                               |
|    |          |                   |         "calls": [                                 |
|    |          |                   |             {                                      |
|    |          |                   |                 "type": "get",                     |
|    |          |                   |                 "scope_values": {                  |
|    |          |                   |                     "table": "hello_get",          |
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
+----+----------+-------------------+----------------------------------------------------+
```

### 3

```
select id, greeting, json, jsonb_pretty(_ctx) as _ctx from hello_get where id in (1,2,17)
```

```
+----+----------+-------------------+----------------------------------------------------+
| id | greeting | json              | _ctx                                               |
+----+----------+-------------------+----------------------------------------------------+
| 2  | Hello    | {"hello":"world"} | {                                                  |
|    |          |                   |     "steampipe": {                                 |
|    |          |                   |         "sdk_version": "5.8.0"                     |
|    |          |                   |     },                                             |
|    |          |                   |     "diagnostics": {                               |
|    |          |                   |         "calls": [                                 |
|    |          |                   |             {                                      |
|    |          |                   |                 "type": "get",                     |
|    |          |                   |                 "scope_values": {                  |
|    |          |                   |                     "table": "hello_get",          |
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
| 1  | Hello    | {"hello":"world"} | {                                                  |
... more data ...
+----+----------+-------------------+----------------------------------------------------+
```

