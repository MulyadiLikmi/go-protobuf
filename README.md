# Generate file
You can generate product.pb.go and product_grpc.pb.go in proto folder by running following command:
```
protoc --go_out=. --go-grpc_out=. proto/product.proto
```

# gRPC Commands
You can use the following grpcurl commands to interact with the gRPC service:
1. Get All Products
```
grpcurl -d "{}" -plaintext localhost:50051 proto.ProductService/
```

2. Get Product By ID
```
// grpcurl -d "{\"id\": \"1\"}" -plaintext localhost:50051 proto.ProductService/GetProductByID
```

3. Add Product
```
grpcurl -d "{\"id\": \"1\", \"name\": \"sample1\", \"stock\": 10}" -plaintext localhost:50051 proto.ProductService/CreateProduct
```

4. Update Product
```
grpcurl -d "{\"id\": \"1\", \"name\": \"updated1\", \"stock\": 20}" -plaintext localhost:50051 proto.ProductService/UpdateProduct
```

5. Delete Product
```
grpcurl -d "{\"id\": \"1\"}" -plaintext localhost:50051 proto.ProductService/DeleteProduct
```

# Testing 
To run tests using the testify framework, you can run following command:
```
go test -v ./test_testify
```

To run tests using Go's built-in testing, you can run following command:
```
go test -v ./test_builtin
```
# Test Cases
1. List Products: Tests listing all products
2. Get Product By ID: Tests retrieval of a product by ID
3. Create Product: Tests if a product can be created successfully
4. Update Product: Tests updating an existing product
5. Delete Product: Tests deletion of a product by ID

# Test Results (Testify)
```
=== RUN   TestProductService_CreateProduct
--- PASS: TestProductService_CreateProduct (0.00s)
=== RUN   TestProductService_GetProductByID
--- PASS: TestProductService_GetProductByID (0.00s)
=== RUN   TestProductService_UpdateProduct
--- PASS: TestProductService_UpdateProduct (0.00s)
=== RUN   TestProductService_DeleteProduct
--- PASS: TestProductService_DeleteProduct (0.00s)
=== RUN   TestProductService_ListProducts
--- PASS: TestProductService_ListProducts (0.00s)
PASS
ok      go-protobuf/test_testify        0.877s
```

# Test Results (Built-in testing)
```
=== RUN   TestProductService_CreateProduct
--- PASS: TestProductService_CreateProduct (0.00s)
=== RUN   TestProductService_GetProductByID
--- PASS: TestProductService_GetProductByID (0.00s)
=== RUN   TestProductService_UpdateProduct
--- PASS: TestProductService_UpdateProduct (0.00s)
=== RUN   TestProductService_DeleteProduct
--- PASS: TestProductService_DeleteProduct (0.00s)
=== RUN   TestProductService_ListProducts
--- PASS: TestProductService_ListProducts (0.00s)
PASS
ok      go-protobuf/test_builtin        0.433s
```