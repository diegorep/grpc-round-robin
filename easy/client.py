import grpc

# import the generated classes
import ones_pb2
import ones_pb2_grpc

# open a gRPC channel
channel = grpc.insecure_channel('localhost:50051')

# create a stub (client)
stub = ones_pb2_grpc.OnesStub(channel)

# create a valid request message
request = ones_pb2.GetOneRequest()
request.value = 1

# make the call
response = stub.EchoOne(request)

# et voila
print("Response type->", type(response))
print("Response value->", response.value)
print("Response's value's type->", type(response.value))
print("Response's value's value->", response.value.value)
