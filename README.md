# rel-in

This is a repo to show a problem with the preloading of rel.  
See this issue: https://github.com/go-rel/rel/issues/283

# How to run
```
docker-compose up -d
go run .
docker-compose down
```

You will see that the select with `LIMIT 999` works, but the one without does not.  
It will panic with `panic: rel: primary key row does not exists` as mysql doesn't allow
`IN` statements with more than 999 values, and therefore the select to preload fails and
then it cannot parse the result propperly.
