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

```
> select *, pg_typeof(json) from hello_list
+----+----------+-------------------+-----------------------------+-----------+
| id | greeting | json              | _ctx                        | pg_typeof |
+----+----------+-------------------+-----------------------------+-----------+
| 1  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | jsonb     |
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | jsonb     |
| 3  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | jsonb     |
+----+----------+-------------------+-----------------------------+-----------+
```

```
> select *, json->>'hello' as json_value from hello_list where id = 2
+----+----------+-------------------+-----------------------------+------------+
| id | greeting | json              | _ctx                        | json_value |
+----+----------+-------------------+-----------------------------+------------+
| 2  | Hello    | {"hello":"world"} | {"connection_name":"hello"} | world      |
+----+----------+-------------------+-----------------------------+------------+
```

```
> select * from hello_list where id = 17
+----+----------+------+------+
| id | greeting | json | _ctx |
+----+----------+------+------+
+----+----------+------+------+
```

