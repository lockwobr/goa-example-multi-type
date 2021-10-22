# goa-example-multi-type

This is an exmample of what I think is a bug. As it is if you build this and hit the endpoint /family `curl --location --request GET 'localhost:9000/family'` you see:

```
{
    "surname": "Harvey",
    "size": 4,
    "kids": [
        {
            "first_name": "Ben",
            "age": 11
        },
        {
            "first_name": "Caroline",
            "age": 16
        }
    ]
}
```

If you remove the struct tags in the design and build again then you get:
(just commented out [line 48](https://github.com/lockwobr/goa-example-multi-type/blob/master/pkg/design/api.go#L48))

```
{
    "surname": "Harvey",
    "size": 4,
    "mother": {
        "first_name": "Katherine",
        "age": 45
    },
    "father": {
        "first_name": "Martin",
        "age": 50
    },
    "kids": [
        {
            "first_name": "Ben",
            "age": 11
        },
        {
            "first_name": "Caroline",
            "age": 16
        }
    ]
}
```