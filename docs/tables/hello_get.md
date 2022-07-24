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
select * from hello_get
```

```
Error: rpc error: code = Internal desc = 'Get' call for table 'hello_get' is missing 1 required qual: column:'id' operator: = (SQLSTATE HV000)
```

### 2

```
select * from hello_get where id = '17'
```

```
+----+----------+-------------------+-----------------------------+
| id | greeting | json              | _ctx                        |
+----+----------+-------------------+-----------------------------+
| 17 | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
+----+----------+-------------------+-----------------------------+
```

### 3

```
select * from hello_get where id in (1,2,17)
```

```
+----+----------+-------------------+-----------------------------+
| id | greeting | json              | _ctx                        |
+----+----------+-------------------+-----------------------------+
| 1  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
| 17 | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
+----+----------+-------------------+-----------------------------+
```

### 4

```
with ids as ( select 1 as id  union select 2  union select 17 )
select 
  *
from 
  hello_get
join 
  ids
using
  (id)
Error: rpc error: code = Internal desc = 'Get' call for table 'hello_get' is missing 1 required qual: column:'id' operator: = (SQLSTATE HV000)
```