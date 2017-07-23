echo "Building protocols"
protoc --go_out=plugins=grpc,import_path=./signalproto:. *.proto
echo "Done"
