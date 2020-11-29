# Get call to home
curl -i -X GET http://localhost:4000/

# Post call to create
curl -i -X POST http://localhost:4000/snippet/create

# Put call fails
curl -i -X PUT http://localhost:4000/snippet/create

# Call with Query
curl -i -X  GET  http://localhost:4000/snippet?id=123
