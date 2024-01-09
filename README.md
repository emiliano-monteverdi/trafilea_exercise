# Trafilea


Save a number into a collection
```
curl --location 'http://localhost:8000/numbers' \
--header 'Content-Type: application/json' \
--data '{
"number":15
}'
```

Retrieve the value of a specific number
```
curl --location 'http://localhost:8000/numbers/15'
```

Retrieve all the numbers in the collection
```
curl --location 'http://localhost:8000/numbers/bulk/value'
```

Retrieve all the numbers type in the collection
```
curl --location 'http://localhost:8000/numbers/bulk/type'
```
