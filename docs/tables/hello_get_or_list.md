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

These examples mix `getGreeting` and  `listGreeting`.

### 1 uses listGreeting

```
select * from hello_get_or_list 
+----+----------+-------------------+-----------------------------+
| id | greeting | json              | _ctx                        |
+----+----------+-------------------+-----------------------------+
| 1  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
| 3  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
+----+----------+-------------------+-----------------------------+
```

### 2 uses getGreeting

```
select * from hello_get_or_list where id = '17'
```

```
+----+----------+-------------------+-----------------------------+
| id | greeting | json              | _ctx                        |
+----+----------+-------------------+-----------------------------+
| 17  | Hello    | {"hello":"world"} | {"connection_name":"hello"} |
+----+----------+-------------------+-----------------------------+
```


### 3 uses getGreeting

```
select * from hello_get_or_list where id in (1,2,17)
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

Message from the plugin SDK:

> executeGetCall - single qual, qual value is a list - executing get for each qual value item, qualValueList: values:{int64_value:1}  values:{int64_value:2}  values:{int64_value:17}

### 4 uses listGreeting

`listGreeting` runs to competion, returning ids 1, 2, and 3, then they're filtered down to 1, 2

```
with ids as ( select 1 as id union select 2 union select 17 )
select 
  *
from 
  hello_get_or_list
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

