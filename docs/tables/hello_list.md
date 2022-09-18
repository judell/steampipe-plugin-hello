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
select *, pg_typeof(json) from hello_list
```

```
+----+----------+-------------------+-----------------------------+-----------+
| id | greeting | json              | _ctx                        | pg_typeof |
+----+----------+-------------------+-----------------------------+-----------+
| 1  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | jsonb     |
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | jsonb     |
| 3  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | jsonb     |
+----+----------+-------------------+-----------------------------+-----------+
```

###  by

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered by Postgres down to 1 row.

```
select *, json->>'hello' as json_value from hello_list where id = 2
```

```
+----+----------+-------------------+-----------------------------+------------+
| id | greeting | json              | _ctx                        | json_value |
+----+----------+-------------------+-----------------------------+------------+
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | world      |
+----+----------+-------------------+-----------------------------+------------+
```

### 3

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered down to 0 rows.

```
select * from hello_list where id = 17
```

```
+----+----------+------+------+
| id | greeting | json | _ctx |
+----+----------+------+------+
+----+----------+------+------+
```

### 4

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered down to 2 rows.

```
select * from hello_list where id in (1,2,17)
```

```
+----+----------+-------------------+-----------------------------+
| id | greeting | json              | _ctx                        |
+----+----------+-------------------+-----------------------------+
| 1  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
+----+----------+-------------------+-----------------------------+
```

### 5

`listGreeting` runs to completion, returning 3 rows, and then the results are filtered down to 2 rows.

```
with ids as ( select 1 as id union select 2 union select 17 )
select 
  *
from 
  hello_list
join 
  ids
using
  (id)
```

```
+----+----------+-------------------+-----------------------------+
| id | greeting | json              | _ctx                        |
+----+----------+-------------------+-----------------------------+
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
| 1  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
+----+----------+-------------------+-----------------------------+
```
